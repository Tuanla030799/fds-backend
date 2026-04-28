package com.fds.backend.template;

import org.apache.ibatis.annotations.Mapper;
import org.apache.ibatis.annotations.Param;

import java.util.List;
import java.util.UUID;

@Mapper
public interface TemplateMapper {
    List<Template> list(@Param("status") String status, @Param("keyword") String keyword,
                        @Param("limit") int limit, @Param("offset") int offset);
    Template findById(@Param("id") UUID id);
    void create(@Param("id") UUID id, @Param("name") String name, @Param("status") String status,
                @Param("note") String note, @Param("fileId") UUID fileId, @Param("createdBy") UUID createdBy);
    void updateStatus(@Param("id") UUID id, @Param("status") String status, @Param("updatedBy") UUID updatedBy);
    void delete(@Param("id") UUID id, @Param("updatedBy") UUID updatedBy);
}
