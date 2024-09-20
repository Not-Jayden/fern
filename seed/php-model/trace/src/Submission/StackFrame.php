<?php

namespace Seed\Submission;

use Seed\Core\SerializableType;
use Seed\Core\JsonProperty;
use Seed\Core\ArrayType;

class StackFrame extends SerializableType
{
    /**
     * @var string $methodName
     */
    #[JsonProperty("methodName")]
    public string $methodName;

    /**
     * @var int $lineNumber
     */
    #[JsonProperty("lineNumber")]
    public int $lineNumber;

    /**
     * @var array<Scope> $scopes
     */
    #[JsonProperty("scopes"), ArrayType([Scope::class])]
    public array $scopes;

    /**
     * @param array{
     *   methodName: string,
     *   lineNumber: int,
     *   scopes: array<Scope>,
     * } $values
     */
    public function __construct(
        array $values,
    ) {
        $this->methodName = $values['methodName'];
        $this->lineNumber = $values['lineNumber'];
        $this->scopes = $values['scopes'];
    }
}
