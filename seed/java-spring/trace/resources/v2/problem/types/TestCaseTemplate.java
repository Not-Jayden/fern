/**
 * This file was auto-generated by Fern from our API Definition.
 */

package resources.v2.problem.types;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonSetter;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;
import core.ObjectMappers;
import java.lang.Object;
import java.lang.String;
import java.util.Objects;
import org.jetbrains.annotations.NotNull;

@JsonInclude(JsonInclude.Include.NON_ABSENT)
@JsonDeserialize(
    builder = TestCaseTemplate.Builder.class
)
public final class TestCaseTemplate {
  private final TestCaseTemplateId templateId;

  private final String name;

  private final TestCaseImplementation implementation;

  private TestCaseTemplate(TestCaseTemplateId templateId, String name,
      TestCaseImplementation implementation) {
    this.templateId = templateId;
    this.name = name;
    this.implementation = implementation;
  }

  @JsonProperty("templateId")
  public TestCaseTemplateId getTemplateId() {
    return templateId;
  }

  @JsonProperty("name")
  public String getName() {
    return name;
  }

  @JsonProperty("implementation")
  public TestCaseImplementation getImplementation() {
    return implementation;
  }

  @java.lang.Override
  public boolean equals(Object other) {
    if (this == other) return true;
    return other instanceof TestCaseTemplate && equalTo((TestCaseTemplate) other);
  }

  private boolean equalTo(TestCaseTemplate other) {
    return templateId.equals(other.templateId) && name.equals(other.name) && implementation.equals(other.implementation);
  }

  @java.lang.Override
  public int hashCode() {
    return Objects.hash(this.templateId, this.name, this.implementation);
  }

  @java.lang.Override
  public String toString() {
    return ObjectMappers.stringify(this);
  }

  public static TemplateIdStage builder() {
    return new Builder();
  }

  public interface TemplateIdStage {
    NameStage templateId(@NotNull TestCaseTemplateId templateId);

    Builder from(TestCaseTemplate other);
  }

  public interface NameStage {
    ImplementationStage name(@NotNull String name);
  }

  public interface ImplementationStage {
    _FinalStage implementation(@NotNull TestCaseImplementation implementation);
  }

  public interface _FinalStage {
    TestCaseTemplate build();
  }

  @JsonIgnoreProperties(
      ignoreUnknown = true
  )
  public static final class Builder implements TemplateIdStage, NameStage, ImplementationStage, _FinalStage {
    private TestCaseTemplateId templateId;

    private String name;

    private TestCaseImplementation implementation;

    private Builder() {
    }

    @java.lang.Override
    public Builder from(TestCaseTemplate other) {
      templateId(other.getTemplateId());
      name(other.getName());
      implementation(other.getImplementation());
      return this;
    }

    @java.lang.Override
    @JsonSetter("templateId")
    public NameStage templateId(@NotNull TestCaseTemplateId templateId) {
      this.templateId = Objects.requireNonNull(templateId, "templateId must not be null");
      return this;
    }

    @java.lang.Override
    @JsonSetter("name")
    public ImplementationStage name(@NotNull String name) {
      this.name = Objects.requireNonNull(name, "name must not be null");
      return this;
    }

    @java.lang.Override
    @JsonSetter("implementation")
    public _FinalStage implementation(@NotNull TestCaseImplementation implementation) {
      this.implementation = Objects.requireNonNull(implementation, "implementation must not be null");
      return this;
    }

    @java.lang.Override
    public TestCaseTemplate build() {
      return new TestCaseTemplate(templateId, name, implementation);
    }
  }
}
