-- +goose Up
-- +goose StatementBegin
INSERT INTO mail_templates (code, subject, content, layout, created_at, updated_at)
VALUES ('reset_password',
        'Yêu cầu đặt lại mật khẩu.',
        '<td style="padding: 1rem; padding-top: 1.5rem !important; padding-bottom: 1.5rem !important;"> <p> Xin chào {{.Name}},<br /> Bạn đã yêu cầu đặt lại mật khẩu. Hãy truy cập vào đường link phía dưới để đặt lại mật khẩu:<br /> </p> <p><a href="{{.Url}}" style="color: #04af04; font-size: 15px; text-decoration: none;">Đặt lại mật khẩu</a></p> <p>Nếu không phải là bạn đã yêu cầu đặt lại mật khẩu, hãy bỏ qua email này.</p> <p>Trân trọng.</p> <div style="margin-top: 0.25rem; font-weight: bold; font-style: italic; color: #04af04;">Đội ngũ phát triển UTCKetnoi</div> </td>',
        'layout1',
        NOW(), NOW());
-- OIDS are not typically used in PostgreSQL
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM mail_templates;
-- +goose StatementEnd

