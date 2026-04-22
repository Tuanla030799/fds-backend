package com.fds.backend.preset;

import jakarta.validation.constraints.NotBlank;
import jakarta.validation.constraints.NotNull;

import java.util.UUID;

public class PresetDtos {
    public record CreateRequest(
            @NotBlank String name,
            String status,
            String note,
            String tags,
            @NotNull UUID fileId,
            Integer sortOrder
    ) {}
}
