package com.fds.backend.design;

import com.fds.backend.common.ApiResponse;
import jakarta.validation.Valid;
import jakarta.validation.constraints.Max;
import jakarta.validation.constraints.Min;
import org.springframework.web.bind.annotation.*;

import java.util.List;
import java.util.UUID;

@RestController
@RequestMapping("/api")
public class DesignSubmissionController {
    private final DesignSubmissionService service;

    public DesignSubmissionController(DesignSubmissionService service) {
        this.service = service;
    }

    @PostMapping("/design-submissions")
    public ApiResponse<Void> create(@Valid @RequestBody DesignSubmissionDtos.CreateRequest body) {
        service.create(body.fullName(), body.address(), body.phone(), body.note(), body.imageUrl());
        return ApiResponse.ok("Created", null);
    }

    @GetMapping("/admin/design-submissions")
    public ApiResponse<List<DesignSubmission>> list(@RequestParam(required = false) String status,
                                                     @RequestParam(required = false) String keyword,
                                                     @RequestParam(defaultValue = "1") @Min(1) int page,
                                                     @RequestParam(defaultValue = "20") @Min(1) @Max(100) int limit) {
        return ApiResponse.ok("OK", service.list(status, keyword, page, limit));
    }

    @PatchMapping("/admin/design-submissions/{id}/status")
    public ApiResponse<Void> updateStatus(@PathVariable UUID id,
                                          @Valid @RequestBody DesignSubmissionDtos.UpdateStatusRequest body) {
        service.updateStatus(id, body.status());
        return ApiResponse.ok("Updated", null);
    }

    @DeleteMapping("/admin/design-submissions/{id}")
    public ApiResponse<Void> delete(@PathVariable UUID id) {
        service.delete(id);
        return ApiResponse.ok("Deleted", null);
    }
}
