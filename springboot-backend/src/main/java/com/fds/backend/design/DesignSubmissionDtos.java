package com.fds.backend.design;

import jakarta.validation.constraints.NotBlank;

public class DesignSubmissionDtos {
    public record CreateRequest(
            @NotBlank String fullName,
            @NotBlank String address,
            @NotBlank String phone,
            String note,
            @NotBlank String imageUrl
    ) {}

    public record UpdateStatusRequest(@NotBlank String status) {}
}
