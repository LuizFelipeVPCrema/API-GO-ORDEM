CREATE TABLE rituais (
    id INTEGER PRIMARY KEY AUTOINCREMENT,

    codigo TEXT NOT NULL UNIQUE,
    nome TEXT NOT NULL,

    elemento TEXT NOT NULL
        CHECK (
            elemento IN (
                'SANGUE',
                'MORTE',
                'CONHECIMENTO',
                'ENERGIA',
                'MEDO'
            )
        ),

    circulo INTEGER NOT NULL
        CHECK (circulo BETWEEN 1 AND 4),

    execucao TEXT NOT NULL
        CHECK (
            execucao IN (
                'LIVRE',
                'MOVIMENTO',
                'PADRAO',
                'COMPLETA',
                'REACAO',
                'ESPECIAL'
            )
        ),

    alcance TEXT NOT NULL,
    alvo TEXT,
    area TEXT,
    duracao TEXT NOT NULL,
    resistencia TEXT,

    custo_pe_base INTEGER NOT NULL
        CHECK (custo_pe_base >= 0),

    requer_componente INTEGER NOT NULL DEFAULT 1
        CHECK (requer_componente IN (0, 1)),

    descricao_resumida TEXT,

    fonte_regra TEXT NOT NULL,
    versao_regra TEXT,

    pagina_referencia INTEGER
        CHECK (
            pagina_referencia IS NULL
            OR pagina_referencia > 0
        ),

    ativo INTEGER NOT NULL DEFAULT 1
        CHECK (ativo IN (0, 1)),

    criado_em DATETIME NOT NULL
        DEFAULT CURRENT_TIMESTAMP,

    atualizado_em DATETIME NOT NULL
        DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE ritual_aprimoramentos (
    id INTEGER PRIMARY KEY AUTOINCREMENT,

    ritual_id INTEGER NOT NULL,

    tipo TEXT NOT NULL
        CHECK (
            tipo IN (
                'DISCENTE',
                'VERDADEIRO',
                'OUTRO'
            )
        ),

    custo_pe_adicional INTEGER NOT NULL
        CHECK (custo_pe_adicional >= 0),

    nex_minimo INTEGER
        CHECK (
            nex_minimo IS NULL
            OR (
                nex_minimo >= 0
                AND nex_minimo <= 100
            )
        ),

    circulo_minimo INTEGER
        CHECK (
            circulo_minimo IS NULL
            OR circulo_minimo BETWEEN 1 AND 4
        ),

    descricao_resumida TEXT,

    ordem_exibicao INTEGER NOT NULL DEFAULT 0,

    ativo INTEGER NOT NULL DEFAULT 1
        CHECK (ativo IN (0, 1)),

    criado_em DATETIME NOT NULL
        DEFAULT CURRENT_TIMESTAMP,

    atualizado_em DATETIME NOT NULL
        DEFAULT CURRENT_TIMESTAMP,

    UNIQUE (
        ritual_id,
        tipo
    ),

    FOREIGN KEY (ritual_id)
        REFERENCES rituais(id)
        ON DELETE CASCADE
);

CREATE INDEX idx_rituais_elemento
    ON rituais(elemento);

CREATE INDEX idx_rituais_circulo
    ON rituais(circulo);

CREATE INDEX idx_rituais_elemento_circulo
    ON rituais(elemento, circulo);

CREATE INDEX idx_ritual_aprimoramentos_ritual
    ON ritual_aprimoramentos(ritual_id);