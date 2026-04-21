package com.fds.backend.preset;

import org.apache.ibatis.annotations.Mapper;
import org.apache.ibatis.annotations.Param;

import java.util.List;
import java.util.UUID;

@Mapper
public interface PresetMapper {
    List<Preset> list(@Param("status") String status, @Param("keyword") String keyword,
                      @Param("limit") int limit, @Param("offset") int offset);
    void create(@Param("id") UUID id, @Param("name") String name, @Param("status") String status,
                @Param("note") String note, @Param("tags") String tags, @Param("imageUrl") String imageUrl,
                @Param("sortOrder") int sortOrder);
    void delete(@Param("id") UUID id);
}
