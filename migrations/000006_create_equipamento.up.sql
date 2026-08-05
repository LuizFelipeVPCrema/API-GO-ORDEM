CREATE TABLE equipamentos (
    id INTEGER PRIMARY KEY AUTOINCREMENT,

    codigo TEXT NOT NULL UNIQUE,
    nome TEXT NOT NULL,

    tipo TEXT NOT NULL
        CHECK (
            tipo IN (
                'ARMA',
                'PROTECAO',
                'ACESSORIO',
                'MUNICAO',
                'ITEM_OPERACIONAL',
                'ITEM_PARANORMAL',
                'OUTRO'
            )
        ),

    categoria_base INTEGER NOT NULL
        CHECK (categoria_base BETWEEN 0 AND 4),

    espacos_base INTEGER NOT NULL
        CHECK (espacos_base >= 0),

    descricao_resumida TEXT,

    fonte_regra TEXT NOT NULL,
    versao_regra TEXT,

    pagina_referencia INTEGER
        CHECK (
            pagina_referencia IS NULL
            OR pagina_referencia > 0
        ),

    ativa INTEGER NOT NULL DEFAULT 1
        CHECK (ativa IN (0, 1)),

    criado_em DATETIME NOT NULL
        DEFAULT CURRENT_TIMESTAMP,

    atualizado_em DATETIME NOT NULL
        DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE modificacoes_equipamento (
    id INTEGER PRIMARY KEY AUTOINCREMENT,

    codigo TEXT NOT NULL UNIQUE,
    nome TEXT NOT NULL,

    incremento_categoria INTEGER NOT NULL
        DEFAULT 0,

    incremento_espacos INTEGER NOT NULL
        DEFAULT 0,

    limite_por_item INTEGER NOT NULL
        DEFAULT 1
        CHECK (limite_por_item > 0),

    cumulativa INTEGER NOT NULL
        DEFAULT 0
        CHECK (cumulativa IN (0, 1)),

    descricao_resumida TEXT,

    fonte_regra TEXT NOT NULL,
    versao_regra TEXT,

    pagina_referencia INTEGER
        CHECK (
            pagina_referencia IS NULL
            OR pagina_referencia > 0
        ),

    ativa INTEGER NOT NULL DEFAULT 1
        CHECK (ativa IN (0, 1)),

    criado_em DATETIME NOT NULL
        DEFAULT CURRENT_TIMESTAMP,

    atualizado_em DATETIME NOT NULL
        DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE modificacao_tipos_equipamento (
    modificacao_id INTEGER NOT NULL,

    tipo_equipamento TEXT NOT NULL
        CHECK (
            tipo_equipamento IN (
                'ARMA',
                'PROTECAO',
                'ACESSORIO',
                'MUNICAO',
                'ITEM_OPERACIONAL',
                'ITEM_PARANORMAL',
                'OUTRO',
                'TODOS'
            )
        ),

    PRIMARY KEY (
        modificacao_id,
        tipo_equipamento
    ),

    FOREIGN KEY (modificacao_id)
        REFERENCES modificacoes_equipamento(id)
        ON DELETE CASCADE
);

CREATE INDEX idx_equipamentos_tipo
    ON equipamentos(tipo);

CREATE INDEX idx_equipamentos_categoria
    ON equipamentos(categoria_base);

CREATE INDEX idx_equipamentos_tipo_categoria
    ON equipamentos(
        tipo,
        categoria_base
    );

CREATE INDEX idx_modificacao_tipos_tipo
    ON modificacao_tipos_equipamento(
        tipo_equipamento
    );