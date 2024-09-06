/**
 * This file was auto-generated by Fern from our API Definition.
 */
package com.seed.multiLineDocs.model.user;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonSetter;
import com.fasterxml.jackson.annotation.Nulls;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;
import com.seed.multiLineDocs.core.ObjectMappers;
import java.util.Objects;
import java.util.Optional;

@JsonInclude(JsonInclude.Include.NON_ABSENT)
@JsonDeserialize(builder = User.Builder.class)
public final class User {
    private final String id;

    private final String name;

    private final Optional<Integer> age;

    private User(String id, String name, Optional<Integer> age) {
        this.id = id;
        this.name = name;
        this.age = age;
    }

    @JsonProperty("id")
    public String getId() {
        return id;
    }

    /**
     * @return The user's name. This name is unique to each user. A few examples are included below:
     * <ul>
     * <li>Alice</li>
     * <li>Bob</li>
     * <li>Charlie</li>
     * </ul>
     */
    @JsonProperty("name")
    public String getName() {
        return name;
    }

    /**
     * @return The user's age.
     */
    @JsonProperty("age")
    public Optional<Integer> getAge() {
        return age;
    }

    @java.lang.Override
    public boolean equals(Object other) {
        if (this == other) return true;
        return other instanceof User && equalTo((User) other);
    }

    private boolean equalTo(User other) {
        return id.equals(other.id) && name.equals(other.name) && age.equals(other.age);
    }

    @java.lang.Override
    public int hashCode() {
        return Objects.hash(this.id, this.name, this.age);
    }

    @java.lang.Override
    public String toString() {
        return ObjectMappers.stringify(this);
    }

    public static IdStage builder() {
        return new Builder();
    }

    public interface IdStage {
        NameStage id(String id);

        Builder from(User other);
    }

    public interface NameStage {
        _FinalStage name(String name);
    }

    public interface _FinalStage {
        User build();

        _FinalStage age(Optional<Integer> age);

        _FinalStage age(Integer age);
    }

    @JsonIgnoreProperties(ignoreUnknown = true)
    public static final class Builder implements IdStage, NameStage, _FinalStage {
        private String id;

        private String name;

        private Optional<Integer> age = Optional.empty();

        private Builder() {}

        @java.lang.Override
        public Builder from(User other) {
            id(other.getId());
            name(other.getName());
            age(other.getAge());
            return this;
        }

        @java.lang.Override
        @JsonSetter("id")
        public NameStage id(String id) {
            this.id = Objects.requireNonNull(id, "id must not be null");
            return this;
        }

        /**
         * <p>The user's name. This name is unique to each user. A few examples are included below:</p>
         * <ul>
         * <li>Alice</li>
         * <li>Bob</li>
         * <li>Charlie</li>
         * </ul>
         * @return Reference to {@code this} so that method calls can be chained together.
         */
        @java.lang.Override
        @JsonSetter("name")
        public _FinalStage name(String name) {
            this.name = Objects.requireNonNull(name, "name must not be null");
            return this;
        }

        /**
         * <p>The user's age.</p>
         * @return Reference to {@code this} so that method calls can be chained together.
         */
        @java.lang.Override
        public _FinalStage age(Integer age) {
            this.age = Optional.ofNullable(age);
            return this;
        }

        @java.lang.Override
        @JsonSetter(value = "age", nulls = Nulls.SKIP)
        public _FinalStage age(Optional<Integer> age) {
            this.age = age;
            return this;
        }

        @java.lang.Override
        public User build() {
            return new User(id, name, age);
        }
    }
}
