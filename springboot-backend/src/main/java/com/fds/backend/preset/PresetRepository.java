package com.fds.backend.preset;

import java.util.List;
import java.util.UUID;

public interface PresetRepository {
    List<Preset> list(String status, String keyword, int limit, int offset);
    void create(UUID id, String name, String status, String note, String tags, String imageUrl, int sortOrder);
    void delete(UUID id);
}
