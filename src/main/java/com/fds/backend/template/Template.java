package com.fds.backend.template;

import java.time.OffsetDateTime;
import java.util.UUID;

public record Template(UUID id, String name, String status, String note, UUID fileId, String imageUrl,
                       OffsetDateTime createdAt) {}
