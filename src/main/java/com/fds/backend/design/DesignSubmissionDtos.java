package com.fds.backend.design;

import jakarta.validation.constraints.NotBlank;
import jakarta.validation.constraints.NotNull;

import java.util.UUID;

public class DesignSubmissionDtos {
    public record CreateRequest(
            @NotBlank String fullName,
            @NotBlank String address,
            @NotBlank String phone,
            String note,
            @NotNull UUID fileId
    ) {}

    public record UpdateStatusRequest(@NotBlank String status) {}
}
