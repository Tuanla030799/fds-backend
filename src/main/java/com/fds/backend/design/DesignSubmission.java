package com.fds.backend.design;

import java.time.OffsetDateTime;
import java.util.UUID;

public record DesignSubmission(UUID id, String fullName, String address, String phone, String note,
                               UUID fileId, String imageUrl, String status, OffsetDateTime createdAt) {}
