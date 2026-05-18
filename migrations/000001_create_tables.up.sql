-- ========================================
-- LIMPIEZA (opcional)
-- ========================================
DROP TABLE IF EXISTS stacks;
DROP TABLE IF EXISTS segment_purpose;
DROP TABLE IF EXISTS purposes;
DROP TABLE IF EXISTS segments;

-- ========================================
-- TABLA: segments
-- ========================================
CREATE TABLE segments
(
    id            INT AUTO_INCREMENT PRIMARY KEY,
    name          VARCHAR(100) NOT NULL,
    is_productive TINYINT(1) NOT NULL DEFAULT 0
);

-- ========================================
-- TABLA: purposes
-- ========================================
CREATE TABLE purposes
(
    id            INT AUTO_INCREMENT PRIMARY KEY,
    name          VARCHAR(100) NOT NULL,
    description   VARCHAR(200),
    is_productive TINYINT(1) NOT NULL DEFAULT 0
);

-- ========================================
-- TABLA: segment_purpose
-- ========================================
CREATE TABLE segment_purpose
(
    segment_id INT NOT NULL,
    purpose_id INT NOT NULL,

    PRIMARY KEY (segment_id, purpose_id),

    CONSTRAINT fk_sp_segment
        FOREIGN KEY (segment_id) REFERENCES segments (id)
            ON DELETE CASCADE,

    CONSTRAINT fk_sp_purpose
        FOREIGN KEY (purpose_id) REFERENCES purposes (id)
            ON DELETE CASCADE
);

-- ========================================
-- TABLA: stacks
-- ========================================
CREATE TABLE stacks
(
    id               INT AUTO_INCREMENT PRIMARY KEY,
    name             VARCHAR(150) NOT NULL,
    description      VARCHAR(200),

    segment_id       INT          NOT NULL,
    purpose_id       INT          NOT NULL,

    application_name VARCHAR(150),

    created_at       DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at       DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    created_by       VARCHAR(100),
    updated_by       VARCHAR(100),

    CONSTRAINT fk_stack_segment
        FOREIGN KEY (segment_id) REFERENCES segments (id),

    CONSTRAINT fk_stack_purpose
        FOREIGN KEY (purpose_id) REFERENCES purposes (id),
);

-- ========================================
-- ÍNDICES
-- ========================================
CREATE INDEX idx_stacks_segment ON stacks (segment_id);
CREATE INDEX idx_stacks_purpose ON stacks (purpose_id);



ALTER TABLE segments
    ADD CONSTRAINT chk_segments_productive CHECK (is_productive IN (0, 1));

ALTER TABLE purposes
    ADD CONSTRAINT chk_purposes_productive CHECK (is_productive IN (0, 1));