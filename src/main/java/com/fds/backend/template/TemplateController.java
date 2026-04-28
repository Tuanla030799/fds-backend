package com.fds.backend.template;

import com.fds.backend.common.ApiResponse;
import jakarta.validation.Valid;
import jakarta.validation.constraints.Max;
import jakarta.validation.constraints.Min;
import org.springframework.web.bind.annotation.*;

import java.util.List;
import java.util.UUID;

@RestController
@RequestMapping("/api")
public class TemplateController {
    private final TemplateService templateService;

    public TemplateController(TemplateService templateService) {
        this.templateService = templateService;
    }

    @GetMapping("/templates")
    public ApiResponse<List<Template>> listPublic(@RequestParam(required = false) String status,
                                                @RequestParam(required = false) String keyword,
                                                @RequestParam(defaultValue = "1") @Min(1) int page,
                                                @RequestParam(defaultValue = "20") @Min(1) @Max(100) int limit) {
        return ApiResponse.ok("OK", templateService.listPublic(keyword, page, limit));
    }

    @GetMapping("/admin/templates")
    public ApiResponse<List<Template>> listAdmin(@RequestParam(required = false) String status,
                                               @RequestParam(required = false) String keyword,
                                               @RequestParam(defaultValue = "1") @Min(1) int page,
                                               @RequestParam(defaultValue = "20") @Min(1) @Max(100) int limit) {
        return ApiResponse.ok("OK", templateService.listAdmin(status, keyword, page, limit));
    }

    @PostMapping("/admin/templates")
    public ApiResponse<Void> create(@Valid @RequestBody TemplateDtos.CreateRequest body) {
        templateService.create(new TemplateService.TemplateCreateRequest(
                body.name(),
                body.status(),
                body.note(),
                body.fileId()
        ));
        return ApiResponse.ok("Created", null);
    }

    @PatchMapping("/admin/templates/{id}/status")
    public ApiResponse<Void> updateStatus(@PathVariable UUID id,
                                          @Valid @RequestBody TemplateDtos.UpdateStatusRequest body) {
        templateService.updateStatus(id, body.status());
        return ApiResponse.ok("Updated", null);
    }

    @DeleteMapping("/admin/templates/{id}")
    public ApiResponse<Void> delete(@PathVariable UUID id) {
        templateService.delete(id);
        return ApiResponse.ok("Deleted", null);
    }
}
