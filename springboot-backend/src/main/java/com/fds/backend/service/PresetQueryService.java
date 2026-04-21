package com.fds.backend.service;

import com.fds.backend.domain.preset.PresetSummary;
import com.fds.backend.mapper.PresetQueryMapper;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class PresetQueryService {
    private final PresetQueryMapper presetQueryMapper;

    public PresetQueryService(PresetQueryMapper presetQueryMapper) {
        this.presetQueryMapper = presetQueryMapper;
    }

    public List<PresetSummary> search(String status, String keyword, int page, int size) {
        int safePage = Math.max(page, 1);
        int safeSize = Math.min(Math.max(size, 1), 100);
        int offset = (safePage - 1) * safeSize;
        return presetQueryMapper.searchPresets(status, keyword, safeSize, offset);
    }
}
