package com.fds.backend.design;

import com.fds.backend.common.ApiResponse;
import jakarta.validation.constraints.Max;
import jakarta.validation.constraints.Min;
import jakarta.validation.constraints.NotBlank;
import org.springframework.web.bind.annotation.*;

import java.util.List;
import java.util.UUID;

@RestController
public class DesignSubmissionController {
    private final DesignSubmissionMapper mapper;

    public DesignSubmissionController(DesignSubmissionMapper mapper) {
        this.mapper = mapper;
    }

    @PostMapping("/api/design-submissions")
    public ApiResponse<Void> create(@RequestBody CreateBody body) {
        mapper.create(UUID.randomUUID(), body.fullName, body.address, body.phone, body.note, body.imageUrl);
        return ApiResponse.ok("Created", null);
    }

    @GetMapping("/api/admin/design-submissions")
    public ApiResponse<List<DesignSubmission>> list(@RequestParam(required = false) String status,
                                                     @RequestParam(required = false) String keyword,
                                                     @RequestParam(defaultValue = "1") @Min(1) int page,
                                                     @RequestParam(defaultValue = "20") @Min(1) @Max(100) int limit) {
        int safeLimit = Math.min(Math.max(limit, 1), 100);
        int offset = (Math.max(page, 1) - 1) * safeLimit;
        return ApiResponse.ok("OK", mapper.list(status, keyword, safeLimit, offset));
    }

    @PatchMapping("/api/admin/design-submissions/{id}/status")
    public ApiResponse<Void> updateStatus(@PathVariable UUID id, @RequestBody UpdateStatusBody body) {
        mapper.updateStatus(id, body.status);
        return ApiResponse.ok("Updated", null);
    }

    @DeleteMapping("/api/admin/design-submissions/{id}")
    public ApiResponse<Void> delete(@PathVariable UUID id) {
        mapper.delete(id);
        return ApiResponse.ok("Deleted", null);
    }

    public static class CreateBody {
        @NotBlank public String fullName;
        @NotBlank public String address;
        @NotBlank public String phone;
        public String note;
        @NotBlank public String imageUrl;
    }

    public static class UpdateStatusBody {
        @NotBlank public String status;
    }
}
