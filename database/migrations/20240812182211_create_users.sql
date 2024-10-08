-- +goose Up
-- +goose StatementBegin
CREATE TABLE `users` (
    `id` int NOT NULL AUTO_INCREMENT,
    `email` varchar(255) NOT NULL,
    `password` text NOT NULL,
    `full_name` varchar(255) NOT NULL,
    `phone_number` varchar(15) NOT NULL,
    `is_active` tinyint(1) NOT NULL,
    `created_at` timestamp NOT NULL,
    `updated_at` timestamp DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `users`;
-- +goose StatementEnd
