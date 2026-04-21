package com.fds.backend.design;

import org.apache.ibatis.annotations.Mapper;
import org.apache.ibatis.annotations.Param;

import java.util.List;
import java.util.UUID;

@Mapper
public interface DesignSubmissionMapper {
    void create(@Param("id") UUID id, @Param("fullName") String fullName, @Param("address") String address,
                @Param("phone") String phone, @Param("note") String note, @Param("imageUrl") String imageUrl);
    List<DesignSubmission> list(@Param("status") String status, @Param("keyword") String keyword,
                                @Param("limit") int limit, @Param("offset") int offset);
    void updateStatus(@Param("id") UUID id, @Param("status") String status);
    void delete(@Param("id") UUID id);
}
