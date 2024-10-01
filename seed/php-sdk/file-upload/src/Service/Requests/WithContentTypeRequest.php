<?php

namespace Seed\Service\Requests;

use Seed\Core\SerializableType;
use Seed\Utils\File;
use Seed\Core\JsonProperty;
use Seed\Service\Types\MyObject;

class WithContentTypeRequest extends SerializableType
{
    /**
     * @var File $file
     */
    public File $file;

    /**
     * @var string $foo
     */
    #[JsonProperty('foo')]
    public string $foo;

    /**
     * @var MyObject $bar
     */
    #[JsonProperty('bar')]
    public MyObject $bar;

    /**
     * @param array{
     *   file: File,
     *   foo: string,
     *   bar: MyObject,
     * } $values
     */
    public function __construct(
        array $values,
    ) {
        $this->file = $values['file'];
        $this->foo = $values['foo'];
        $this->bar = $values['bar'];
    }
}
