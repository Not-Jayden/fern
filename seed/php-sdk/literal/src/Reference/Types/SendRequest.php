<?php

namespace Seed\Reference\Types;

use Seed\Core\Json\JsonSerializableType;
use Seed\Core\Json\JsonProperty;

class SendRequest extends JsonSerializableType
{
    /**
     * @var string $prompt
     */
    #[JsonProperty('prompt')]
    public string $prompt;

    /**
     * @var string $query
     */
    #[JsonProperty('query')]
    public string $query;

    /**
     * @var bool $stream
     */
    #[JsonProperty('stream')]
    public bool $stream;

    /**
     * @var string $context
     */
    #[JsonProperty('context')]
    public string $context;

    /**
     * @var ?string $maybeContext
     */
    #[JsonProperty('maybeContext')]
    public ?string $maybeContext;

    /**
     * @var ContainerObject $containerObject
     */
    #[JsonProperty('containerObject')]
    public ContainerObject $containerObject;

    /**
     * @param array{
     *   prompt: string,
     *   query: string,
     *   stream: bool,
     *   context: string,
     *   maybeContext?: ?string,
     *   containerObject: ContainerObject,
     * } $values
     */
    public function __construct(
        array $values,
    ) {
        $this->prompt = $values['prompt'];
        $this->query = $values['query'];
        $this->stream = $values['stream'];
        $this->context = $values['context'];
        $this->maybeContext = $values['maybeContext'] ?? null;
        $this->containerObject = $values['containerObject'];
    }
}
