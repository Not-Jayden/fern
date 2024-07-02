from ast import arg
from email.mime import base
from typing import List, Optional, Set

from fern_python.codegen.source_file import ConditionalTree, IfConditionLeaf

import fern.ir.resources as ir_types
from typing_extensions import Never

from fern_python.codegen import AST, LocalClassReference, SourceFile
from fern_python.external_dependencies import Pydantic
from fern_python.pydantic_codegen import PydanticField, PydanticModel

from ....context import PydanticGeneratorContext
from ...custom_config import PydanticModelCustomConfig
from ...fern_aware_pydantic_model import FernAwarePydanticModel
from ..abstract_type_generator import AbstractTypeGenerator
from ..get_visit_method import VisitableItem, VisitorArgument, get_visit_method

VISITOR_RETURN_TYPE = AST.GenericTypeVar(name="T_Result")
BUILDER_ARGUMENT_NAME = "value"


class DiscriminatedUnionWithUtilsGenerator(AbstractTypeGenerator):
    def __init__(
        self,
        name: ir_types.DeclaredTypeName,
        union: ir_types.UnionTypeDeclaration,
        context: PydanticGeneratorContext,
        source_file: SourceFile,
        custom_config: PydanticModelCustomConfig,
        docs: Optional[str],
        snippet: Optional[str] = None,
    ):
        super().__init__(
            context=context, custom_config=custom_config, source_file=source_file, docs=docs, snippet=snippet
        )
        self._name = name
        self._union = union


    def _generate_conditional_base_class(self, conditional_class_name: str, internal_single_union_types: List[LocalClassReference]) -> LocalClassReference:
        if self._custom_config.skip_validation:
            root_type = AST.TypeHint.annotated(
                type=AST.TypeHint.union(
                    *(
                        AST.TypeHint(type=internal_single_union_type)
                        for internal_single_union_type in internal_single_union_types
                    ),
                ),
                annotation=AST.Expression(
                    AST.ClassInstantiation(
                        class_=self._context.core_utilities.get_union_metadata(),
                        kwargs=[
                            (
                                "discriminant",
                                AST.Expression(f'"{self._union.discriminant.wire_value}"'),
                            )
                        ],
                    )
                ),
            )
        else:
            root_type = AST.TypeHint.union(
                *(
                    AST.TypeHint(type=internal_single_union_type)
                    for internal_single_union_type in internal_single_union_types
                ),
            )

        root_type_annotation = AST.Expression(
            AST.FunctionInvocation(
                function_definition=Pydantic.Field(),
                kwargs=[
                    (
                        "discriminator",
                        AST.Expression(
                            f'"{self._get_discriminant_attr_name()}"',
                        ),
                    )
                ],
            )
        ) if len(internal_single_union_types) != 1 else None

        annotated_type_hint = AST.TypeHint.annotated(
            type=root_type,
            annotation=root_type_annotation,
        ) if root_type_annotation is not None else root_type
        
        conditional_class_declaration_v2 = AST.ClassDeclaration(name=conditional_class_name)
        conditional_class_declaration_v2.add_statement(
            AST.VariableDeclaration(name="root", type_hint=annotated_type_hint)
        )
        conditional_class_declaration_v2.add_method(
                AST.FunctionDeclaration(
                    name="get_as_union",
                    signature=AST.FunctionSignature(
                        return_type=root_type,
                    ),
                    body=AST.CodeWriter("return self.root"),
                )
            )

        conditional_class_declaration_v1 = AST.ClassDeclaration(name=conditional_class_name)
        conditional_class_declaration_v1.add_statement(
            AST.VariableDeclaration(name="root", type_hint=annotated_type_hint)
        )
        conditional_class_declaration_v1.add_method(
                AST.FunctionDeclaration(
                    name="get_as_union",
                    signature=AST.FunctionSignature(
                        return_type=root_type,
                    ),
                    body=AST.CodeWriter("return self.__root__"),
                )
            )

        conditional_class = self._source_file.add_conditional_class_declaration(
            declaration=conditional_class_declaration_v2,
            conditional_tree=ConditionalTree(
                conditions=[
                    IfConditionLeaf(
                        condition=AST.Expression(self._context.core_utilities.get_is_pydantic_v2()),
                        code=conditional_class_declaration_v2
                    )
                ],
                else_code=conditional_class_declaration_v2
            )
        )

        return conditional_class

    def generate(self) -> None:
        factory_declaration = AST.ClassDeclaration(name="_Factory")
        factory = self._source_file.add_class_declaration(factory_declaration)

        model_name = self._context.get_class_name_for_type_id(self._name.type_id)
        conditional_class_name = f"_{model_name}Base"
        dummy_base_class_declaration = AST.ClassDeclaration(name=conditional_class_name)
        dummy_base_class = self._source_file.get_dummy_class_declaration(dummy_base_class_declaration)
        
        with FernAwarePydanticModel(
            class_name=model_name,
            type_name=self._name,
            context=self._context,
            custom_config=self._custom_config,
            source_file=self._source_file,
            docstring=self._docs,
            snippet=self._snippet,
            base_models=[dummy_base_class],
        ) as external_pydantic_model:
            external_pydantic_model.add_class_var_unsafe(
                name="factory",
                type_hint=AST.TypeHint(type=factory),
                initializer=AST.Expression(AST.ClassInstantiation(class_=factory)),
            )

            internal_single_union_types: List[LocalClassReference] = []

            internal_union = self._source_file.add_class_declaration(
                declaration=AST.ClassDeclaration(name="_" + external_pydantic_model.get_class_name()),
            )

            for single_union_type in self._union.types:
                with PydanticModel(
                    version=self._custom_config.version,
                    name=single_union_type.discriminant_value.name.pascal_case.safe_name,
                    source_file=self._source_file,
                    base_models=single_union_type.shape.visit(
                        same_properties_as_object=lambda type_name: [
                            self._context.get_class_reference_for_type_id(type_name.type_id)
                        ],
                        single_property=lambda _: [],
                        no_properties=lambda: [],
                    ),
                    parent=internal_union,
                    frozen=self._custom_config.frozen,
                    orm_mode=self._custom_config.orm_mode,
                    smart_union=self._custom_config.smart_union,
                    pydantic_base_model=self._context.core_utilities.get_unchecked_pydantic_base_model(),
                    require_optional_fields=self._custom_config.require_optional_fields,
                    is_pydantic_v2=self._context.core_utilities.get_is_pydantic_v2(),
                    universal_field_validator=self._context.core_utilities.universal_field_validator,
                    universal_root_validator=self._context.core_utilities.universal_root_validator,
                ) as internal_pydantic_model_for_single_union_type:
                    internal_single_union_type = internal_pydantic_model_for_single_union_type.to_reference()
                    internal_single_union_types.append(internal_single_union_type)

                    discriminant_field = self._get_discriminant_field_for_single_union_type(
                        single_union_type=single_union_type
                    )

                    internal_pydantic_model_for_single_union_type.add_field(discriminant_field)

                    shape = single_union_type.shape.get_as_union()
                    if shape.properties_type == "singleProperty":
                        internal_pydantic_model_for_single_union_type.add_field(
                            PydanticField(
                                name=shape.name.name.snake_case.safe_name,
                                pascal_case_field_name=shape.name.name.pascal_case.safe_name,
                                json_field_name=shape.name.wire_value,
                                type_hint=self._context.get_type_hint_for_type_reference(type_reference=shape.type),
                            )
                        )

                    factory_declaration.add_method(
                        AST.FunctionDeclaration(
                            name=single_union_type.discriminant_value.name.snake_case.safe_name,
                            signature=AST.FunctionSignature(
                                parameters=single_union_type.shape.visit(
                                    same_properties_as_object=lambda type_name: [
                                        AST.FunctionParameter(
                                            name=BUILDER_ARGUMENT_NAME,
                                            type_hint=self._context.get_type_hint_for_type_reference(
                                                ir_types.TypeReference.factory.named(type_name)
                                            ),
                                        )
                                    ],
                                    single_property=lambda property: [
                                        AST.FunctionParameter(
                                            name=BUILDER_ARGUMENT_NAME,
                                            type_hint=self._context.get_type_hint_for_type_reference(property.type),
                                        )
                                    ],
                                    no_properties=lambda: None,
                                ),
                                return_type=self._context.get_type_hint_for_type_reference(
                                    ir_types.TypeReference.factory.named(self._name)
                                ),
                            ),
                            body=AST.CodeWriter(
                                self._create_body_writer(
                                    single_union_type=single_union_type,
                                    internal_single_union_type=internal_single_union_type,
                                    external_union=external_pydantic_model.to_reference(),
                                )
                            ),
                        ),
                    )

                    # if any of our inherited fields are forward refs, we need to call
                    # update_forwards_refs()

                    # we assume that the forward-refed types are the ones
                    # that circularly reference this union type
                    referenced_type_ids: Set[ir_types.TypeId] = single_union_type.shape.visit(
                        same_properties_as_object=lambda type_name: self._context.get_referenced_types_of_type_declaration(
                            self._context.get_declaration_for_type_id(type_name.type_id),
                        ),
                        single_property=lambda single_property: self._context.get_referenced_types_of_type_reference(
                            single_property.type
                        ),
                        no_properties=lambda: set(),
                    )
                    forward_refed_types = [
                        referenced_type_id
                        for referenced_type_id in referenced_type_ids
                        if self._context.does_type_reference_other_type(referenced_type_id, self._name.type_id)
                    ]

                    if len(forward_refed_types) > 0:
                        # when calling update_forward_refs, Pydantic will throw
                        # if an inherited field's type is not defined in this
                        # file. https://github.com/pydantic/pydantic/issues/4902.
                        # as a workaround, we explicitly pass references to update_forward_refs
                        # so they are in scope
                        internal_pydantic_model_for_single_union_type.update_forward_refs(
                            {self._context.get_class_reference_for_type_id(type_id) for type_id in forward_refed_types}
                        )

                        # to avoid issues with circular dependencies, make sure all imports
                        # that reference this type appear after the main (exported) model for the union.
                        # FernAwarePydanticModel will automatically add the import constraint if the
                        # referenced type_name circularly references this type.
                        for type_id in forward_refed_types:
                            external_pydantic_model.add_ghost_reference(type_id)

            self._generate_conditional_base_class(conditional_class_name, internal_single_union_types)
            external_pydantic_model.add_method_unsafe(
                get_visit_method(
                    items=[
                        VisitableItem(
                            parameter_name=single_union_type.discriminant_value.name.snake_case.safe_name,
                            expected_value=f'"{single_union_type.discriminant_value.wire_value}"',
                            visitor_argument=single_union_type.shape.visit(
                                same_properties_as_object=lambda type_name: VisitorArgument(
                                    expression=AST.Expression(
                                        AST.FunctionInvocation(
                                            function_definition=self._context.get_class_reference_for_type_id(
                                                type_name.type_id
                                            ),
                                            args=[
                                                AST.Expression(
                                                    f'self.get_as_union().dict(exclude_unset=True, exclude={{"{self._union.discriminant.wire_value}"}})',
                                                    spread=AST.ExpressionSpread.TWO_ASTERISKS,
                                                )
                                            ],
                                        )
                                    ),
                                    type=external_pydantic_model.get_type_hint_for_type_reference(
                                        ir_types.TypeReference.factory.named(type_name)
                                    ),
                                ),
                                single_property=lambda property: VisitorArgument(
                                    expression=AST.Expression(
                                        f"self.get_as_union().{get_field_name_for_single_property(property)}"
                                    ),
                                    type=external_pydantic_model.get_type_hint_for_type_reference(property.type),
                                ),
                                no_properties=lambda: None,
                            ),
                        )
                        for single_union_type in self._union.types
                    ],
                    reference_to_current_value=f"self.get_as_union().{self._get_discriminant_attr_name()}",
                )
            )

    def _create_body_writer(
        self,
        single_union_type: ir_types.SingleUnionType,
        internal_single_union_type: AST.ClassReference,
        external_union: AST.ClassReference,
    ) -> AST.CodeWriterFunction:
        def write(writer: AST.NodeWriter) -> None:
            # explicit typing needed to help mypy
            no_expressions: List[AST.Expression] = []

            internal_single_union_type_instantiation = AST.ClassInstantiation(
                class_=internal_single_union_type,
                args=single_union_type.shape.visit(
                    same_properties_as_object=lambda type_name: [
                        AST.Expression(
                            f"{BUILDER_ARGUMENT_NAME}.dict(exclude_unset=True)",
                            spread=AST.ExpressionSpread.TWO_ASTERISKS,
                        )
                    ],
                    single_property=lambda property: no_expressions,
                    no_properties=lambda: no_expressions,
                ),
                kwargs=[
                    (
                        self._get_discriminant_attr_name(),
                        self._get_discriminant_value_for_single_union_type(single_union_type),
                    )
                ]
                + (
                    single_union_type.shape.visit(
                        same_properties_as_object=lambda type_name: [],
                        single_property=lambda property: [
                            (get_field_name_for_single_property(property), AST.Expression(BUILDER_ARGUMENT_NAME))
                        ],
                        no_properties=lambda: [],
                    )
                ),
            )

            sub_union_instantiation = AST.ClassInstantiation(
                class_=external_union,
                args=[AST.Expression(internal_single_union_type_instantiation)],
            )

            writer.write("return ")
            writer.write_node(sub_union_instantiation)

        return write

    def _get_discriminant_field_for_single_union_type(
        self, single_union_type: ir_types.SingleUnionType
    ) -> PydanticField:
        discriminant_value = self._get_discriminant_value_for_single_union_type(single_union_type)
        return PydanticField(
            name=self._get_discriminant_attr_name(),
            pascal_case_field_name=self._union.discriminant.name.pascal_case.safe_name,
            type_hint=AST.TypeHint.literal(discriminant_value),
            json_field_name=self._union.discriminant.wire_value,
            default_value=discriminant_value,
        )

    def _get_discriminant_attr_name(self) -> str:
        return self._union.discriminant.name.snake_case.safe_name

    def _get_discriminant_value_for_single_union_type(
        self,
        single_union_type: ir_types.SingleUnionType,
    ) -> AST.Expression:
        return AST.Expression(f'"{single_union_type.discriminant_value.wire_value}"')


def assert_never(arg: Never) -> Never:
    raise AssertionError("Expected code to be unreachable")


def get_field_name_for_single_property(property: ir_types.SingleUnionTypeProperty) -> str:
    return property.name.name.snake_case.safe_name
