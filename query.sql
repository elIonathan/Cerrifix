-- name: CreateRep :exec
INSERT INTO reparaciones (
    apellido, telefono, equipo, cable, pata, soporte, control, estado, falla, fecha_ingreso, fecha_entrega, costo, observaciones
) VALUES (
    ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
);

-- name: GetRepByID :one
SELECT 
    id, apellido, telefono, equipo, cable, pata, soporte, control, estado, falla, fecha_ingreso, fecha_entrega, costo, observaciones
FROM 
    reparaciones
WHERE 
    id = ?;

-- name: GetRepByApellido :one
SELECT 
    id, apellido, telefono, equipo, cable, pata, soporte, control, estado, falla, fecha_ingreso, fecha_entrega, costo, observaciones
FROM 
    reparaciones
WHERE 
    apellido = ?;

-- name: GetRepByTel :one
SELECT 
    id, apellido, telefono, equipo, cable, pata, soporte, control, estado, falla, fecha_ingreso, fecha_entrega, costo, observaciones
FROM 
    reparaciones
WHERE 
    telefono = ?;

-- name: ListRepaByApellido :many
SELECT 
    id, apellido, telefono, equipo, cable, pata, soporte, control, estado, falla, fecha_ingreso, fecha_entrega, costo, observaciones
FROM 
    reparaciones
WHERE 
    apellido = ?;

-- name: ListRepaByTel :many
SELECT 
    id, apellido, telefono, equipo, cable, pata, soporte, control, estado, falla, fecha_ingreso, fecha_entrega, costo, observaciones
FROM 
    reparaciones
WHERE 
    telefono = ?;

-- name: GetRepByTelParcial :many
SELECT 
    id, apellido, telefono, equipo, cable, pata, soporte, control, estado, falla, fecha_ingreso, fecha_entrega, costo, observaciones
FROM 
    reparaciones
WHERE 
    telefono LIKE '%' || ?;
