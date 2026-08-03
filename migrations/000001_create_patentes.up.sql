CREATE TABLE IF NOT EXISTS patentes (
    id INTEGER PRIMARY KEY AUTOINCREMENT,

    codigo TEXT NOT NULL UNIQUE,
    nome TEXT NOT NULL,

    pontos_prestigio_minimos INTEGER NOT NULL
        CHECK (pontos_prestigio_minimos >= 0),

    limite_credito TEXT NOT NULL
        CHECK (
            limite_credito IN (
                'BAIXO',
                'MEDIO',
                'ALTO',
                'ILIMITADO'
            )
        ),

    nivel_hierarquico INTEGER NOT NULL UNIQUE,

    ativa INTEGER NOT NULL DEFAULT 1
        CHECK (ativa IN (0, 1)),

    criado_em DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    atualizado_em DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS patente_limites_item (
    patente_id INTEGER NOT NULL,

    categoria INTEGER NOT NULL
        CHECK (categoria BETWEEN 1 AND 4),

    quantidade_maxima INTEGER NOT NULL
        CHECK (quantidade_maxima >= 0),

    PRIMARY KEY (
        patente_id,
        categoria
    ),

    FOREIGN KEY (patente_id)
        REFERENCES patentes(id)
        ON DELETE CASCADE
);

INSERT OR IGNORE INTO patentes (
    id,
    codigo,
    nome,
    pontos_prestigio_minimos,
    limite_credito,
    nivel_hierarquico
)
VALUES
    (1, 'RECRUTA', 'Recruta', 0, 'BAIXO', 1),
    (2, 'OPERADOR', 'Operador', 20, 'MEDIO', 2),
    (
        3,
        'AGENTE_ESPECIAL',
        'Agente Especial',
        50,
        'MEDIO',
        3
    ),
    (
        4,
        'OFICIAL_OPERACOES',
        'Oficial de Operações',
        100,
        'ALTO',
        4
    ),
    (
        5,
        'AGENTE_ELITE',
        'Agente de Elite',
        200,
        'ILIMITADO',
        5
    );

INSERT OR IGNORE INTO patente_limites_item (
    patente_id,
    categoria,
    quantidade_maxima
)
VALUES
    (1, 1, 2),
    (1, 2, 0),
    (1, 3, 0),
    (1, 4, 0),

    (2, 1, 3),
    (2, 2, 1),
    (2, 3, 0),
    (2, 4, 0),

    (3, 1, 3),
    (3, 2, 2),
    (3, 3, 1),
    (3, 4, 0),

    (4, 1, 3),
    (4, 2, 3),
    (4, 3, 2),
    (4, 4, 1),

    (5, 1, 3),
    (5, 2, 3),
    (5, 3, 3),
    (5, 4, 2);