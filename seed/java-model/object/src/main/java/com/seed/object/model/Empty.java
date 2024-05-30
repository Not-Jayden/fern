/**
 * This file was auto-generated by Fern from our API Definition.
 */
package com.seed.object.model;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;
import com.seed.object.core.ObjectMappers;

@JsonInclude(JsonInclude.Include.NON_EMPTY)
@JsonDeserialize(builder = Empty.Builder.class)
public final class Empty {
    private Empty() {}

    @java.lang.Override
    public boolean equals(Object other) {
        if (this == other) return true;
        return other instanceof Empty;
    }

    @java.lang.Override
    public String toString() {
        return ObjectMappers.stringify(this);
    }

    public static Builder builder() {
        return new Builder();
    }

    @JsonIgnoreProperties(ignoreUnknown = true)
    public static final class Builder {
        private Builder() {}

        public Builder from(Empty other) {
            return this;
        }

        public Empty build() {
            return new Empty();
        }
    }
}
