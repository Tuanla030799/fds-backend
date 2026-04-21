package com.fds.backend.preset;

import jakarta.validation.constraints.NotBlank;

public class PresetDtos {
    public record CreateRequest(
            @NotBlank String name,
            String status,
            String note,
            String tags,
            @NotBlank String imageUrl,
            Integer sortOrder
    ) {}
}
