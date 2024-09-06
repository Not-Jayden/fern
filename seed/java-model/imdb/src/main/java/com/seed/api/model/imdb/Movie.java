/**
 * This file was auto-generated by Fern from our API Definition.
 */
package com.seed.api.model.imdb;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonSetter;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;
import com.seed.api.core.ObjectMappers;
import java.util.Objects;
import org.jetbrains.annotations.NotNull;

@JsonInclude(JsonInclude.Include.NON_ABSENT)
@JsonDeserialize(builder = Movie.Builder.class)
public final class Movie {
    private final String id;

    private final String title;

    private final double rating;

    private Movie(String id, String title, double rating) {
        this.id = id;
        this.title = title;
        this.rating = rating;
    }

    @JsonProperty("id")
    public String getId() {
        return id;
    }

    @JsonProperty("title")
    public String getTitle() {
        return title;
    }

    /**
     * @return The rating scale is one to five stars
     */
    @JsonProperty("rating")
    public double getRating() {
        return rating;
    }

    @java.lang.Override
    public boolean equals(Object other) {
        if (this == other) return true;
        return other instanceof Movie && equalTo((Movie) other);
    }

    private boolean equalTo(Movie other) {
        return id.equals(other.id) && title.equals(other.title) && rating == other.rating;
    }

    @java.lang.Override
    public int hashCode() {
        return Objects.hash(this.id, this.title, this.rating);
    }

    @java.lang.Override
    public String toString() {
        return ObjectMappers.stringify(this);
    }

    public static IdStage builder() {
        return new Builder();
    }

    public interface IdStage {
        TitleStage id(@NotNull String id);

        Builder from(Movie other);
    }

    public interface TitleStage {
        RatingStage title(@NotNull String title);
    }

    public interface RatingStage {
        _FinalStage rating(double rating);
    }

    public interface _FinalStage {
        Movie build();
    }

    @JsonIgnoreProperties(ignoreUnknown = true)
    public static final class Builder implements IdStage, TitleStage, RatingStage, _FinalStage {
        private String id;

        private String title;

        private double rating;

        private Builder() {}

        @java.lang.Override
        public Builder from(Movie other) {
            id(other.getId());
            title(other.getTitle());
            rating(other.getRating());
            return this;
        }

        @java.lang.Override
        @JsonSetter("id")
        public TitleStage id(@NotNull String id) {
            this.id = Objects.requireNonNull(id, "id must not be null");
            return this;
        }

        @java.lang.Override
        @JsonSetter("title")
        public RatingStage title(@NotNull String title) {
            this.title = Objects.requireNonNull(title, "title must not be null");
            return this;
        }

        /**
         * <p>The rating scale is one to five stars</p>
         * @return Reference to {@code this} so that method calls can be chained together.
         */
        @java.lang.Override
        @JsonSetter("rating")
        public _FinalStage rating(double rating) {
            this.rating = rating;
            return this;
        }

        @java.lang.Override
        public Movie build() {
            return new Movie(id, title, rating);
        }
    }
}
