-- schema.sql
CREATE TABLE reparaciones (
    id INTEGER PRIMARY KEY AUTOINCREMENT, -- Un identificador único para cada registro
    apellido TEXT NOT NULL,
    telefono TEXT NOT NULL,
    equipo TEXT NOT NULL,    -- Campo para especificar el tipo de equipo
    cable INTEGER NOT NULL,  -- 0 o 1 para booleano
    pata INTEGER NOT NULL,   -- 0 o 1 para booleano
    soporte INTEGER NOT NULL, -- 0 o 1 para booleano
    control INTEGER NOT NULL, -- 0 o 1 para booleano
    estado TEXT NOT NULL,     -- Estado del equipo o reparación
    falla TEXT NOT NULL,      -- Descripción de la falla
    fecha_ingreso DATE NOT NULL, -- Fecha en que se recibe el equipo
    fecha_entrega DATE,       -- Fecha estimada o real de entrega (puede ser NULL si no está definida aún)
    costo INTEGER,               -- Costo de la reparación (puede ser NULL si no se ha estimado)
    observaciones TEXT        -- Campo opcional para notas adicionales
);
