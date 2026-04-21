package com.fds.backend.preset;

import java.time.OffsetDateTime;
import java.util.UUID;

public record Preset(UUID id, String name, String status, String note, String tags, String imageUrl,
                     Integer sortOrder, OffsetDateTime createdAt) {}
