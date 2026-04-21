package com.fds.backend.mapper;

import com.fds.backend.domain.preset.PresetSummary;
import org.apache.ibatis.annotations.Mapper;
import org.apache.ibatis.annotations.Param;

import java.util.List;

@Mapper
public interface PresetQueryMapper {
    List<PresetSummary> searchPresets(
            @Param("status") String status,
            @Param("keyword") String keyword,
            @Param("limit") int limit,
            @Param("offset") int offset
    );
}
