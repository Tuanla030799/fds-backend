package com.fds.backend.domain.preset;

import java.time.OffsetDateTime;
import java.util.UUID;

public record PresetSummary(
        UUID id,
        String name,
        String status,
        String imageUrl,
        Integer sortOrder,
        OffsetDateTime createdAt,
        int totalCount
) {}
