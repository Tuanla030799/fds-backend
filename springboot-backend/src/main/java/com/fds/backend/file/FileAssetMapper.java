package com.fds.backend.file;

import org.apache.ibatis.annotations.Mapper;
import org.apache.ibatis.annotations.Param;

import java.util.UUID;

@Mapper
public interface FileAssetMapper {
    void create(@Param("id") UUID id, @Param("path") String path, @Param("folder") String folder,
                @Param("status") String status);
}
