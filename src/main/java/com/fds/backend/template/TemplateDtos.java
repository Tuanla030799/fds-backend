package com.fds.backend.template;

import jakarta.validation.constraints.NotBlank;
import jakarta.validation.constraints.NotNull;

import java.util.UUID;

public class TemplateDtos {
    public record CreateRequest(
            @NotBlank String name,
            String status,
            String note,
            @NotNull UUID fileId
    ) {}

    public record UpdateStatusRequest(@NotBlank String status) {}
}
