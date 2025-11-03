-- Spatial Index Examples for TSQL Parser
-- Testing all variations of CREATE SPATIAL INDEX statements

-- Basic CREATE SPATIAL INDEX statements
CREATE SPATIAL INDEX IX_SpatialBasic_1 ON SpatialTable (GeometryColumn) USING GEOMETRY_AUTO_GRID;
CREATE SPATIAL INDEX IX_SpatialBasic_2 ON SpatialTable (GeometryColumn) USING GEOMETRY_GRID;
CREATE SPATIAL INDEX IX_SpatialBasic_3 ON SpatialTable (GeographyColumn) USING GEOGRAPHY_AUTO_GRID;
CREATE SPATIAL INDEX IX_SpatialBasic_4 ON SpatialTable (GeographyColumn) USING GEOGRAPHY_GRID;

-- Spatial indexes with schema-qualified table names
CREATE SPATIAL INDEX IX_SchemaQualified_1 ON dbo.SpatialTable (LocationColumn) USING GEOMETRY_AUTO_GRID;
CREATE SPATIAL INDEX IX_SchemaQualified_2 ON [Sales].[LocationData] ([GeoPoint]) USING GEOGRAPHY_AUTO_GRID;
CREATE SPATIAL INDEX IX_SchemaQualified_3 ON MyDatabase.dbo.SpatialTable (SpatialColumn) USING GEOMETRY_GRID;

-- Spatial indexes with bounding box (supporting negative coordinates)
CREATE SPATIAL INDEX IX_BoundingBox_1 ON SpatialTable (GeometryColumn)
USING GEOMETRY_GRID
WITH (BOUNDING_BOX = (0, 0, 100, 100));

CREATE SPATIAL INDEX IX_BoundingBox_2 ON SpatialTable (GeometryColumn)
USING GEOMETRY_GRID
WITH (BOUNDING_BOX = (-180, -90, 180, 90));

CREATE SPATIAL INDEX IX_BoundingBox_3 ON SpatialTable (GeometryColumn)
USING GEOMETRY_GRID
WITH (BOUNDING_BOX = (-123.456, -78.901, 234.567, 89.012));

-- Spatial indexes with grid level specifications
CREATE SPATIAL INDEX IX_GridLevels_1 ON SpatialTable (GeometryColumn)
USING GEOMETRY_GRID
WITH (GRIDS = (LEVEL_1 = LOW));

CREATE SPATIAL INDEX IX_GridLevels_2 ON SpatialTable (GeometryColumn)
USING GEOMETRY_GRID
WITH (GRIDS = (LEVEL_1 = HIGH, LEVEL_2 = MEDIUM));

CREATE SPATIAL INDEX IX_GridLevels_3 ON SpatialTable (GeometryColumn)
USING GEOMETRY_GRID
WITH (GRIDS = (LEVEL_1 = HIGH, LEVEL_2 = HIGH, LEVEL_3 = MEDIUM));

CREATE SPATIAL INDEX IX_GridLevels_4 ON SpatialTable (GeometryColumn)
USING GEOMETRY_GRID
WITH (GRIDS = (LEVEL_1 = HIGH, LEVEL_2 = HIGH, LEVEL_3 = MEDIUM, LEVEL_4 = LOW));

-- Spatial indexes with cells per object
CREATE SPATIAL INDEX IX_CellsPerObject_1 ON SpatialTable (GeometryColumn)
USING GEOMETRY_GRID
WITH (CELLS_PER_OBJECT = 1);

CREATE SPATIAL INDEX IX_CellsPerObject_2 ON SpatialTable (GeometryColumn)
USING GEOMETRY_GRID
WITH (CELLS_PER_OBJECT = 16);

CREATE SPATIAL INDEX IX_CellsPerObject_3 ON SpatialTable (GeometryColumn)
USING GEOMETRY_GRID
WITH (CELLS_PER_OBJECT = 8192);

-- Spatial indexes with multiple options
CREATE SPATIAL INDEX IX_MultipleOptions_1 ON SpatialTable (GeometryColumn)
USING GEOMETRY_GRID
WITH (
    BOUNDING_BOX = (0, 0, 1000, 1000),
    GRIDS = (LEVEL_1 = HIGH, LEVEL_2 = MEDIUM),
    CELLS_PER_OBJECT = 32
);

CREATE SPATIAL INDEX IX_MultipleOptions_2 ON SpatialTable (GeometryColumn)
USING GEOMETRY_GRID
WITH (
    BOUNDING_BOX = (-50.5, -25.25, 50.5, 25.25),
    GRIDS = (LEVEL_1 = HIGH, LEVEL_2 = HIGH, LEVEL_3 = MEDIUM, LEVEL_4 = LOW),
    CELLS_PER_OBJECT = 64
);

-- Spatial indexes with traditional index options
CREATE SPATIAL INDEX IX_TraditionalOptions_1 ON SpatialTable (GeometryColumn)
USING GEOMETRY_GRID
WITH (
    BOUNDING_BOX = (0, 0, 100, 100),
    PAD_INDEX = ON,
    FILLFACTOR = 90
);

CREATE SPATIAL INDEX IX_TraditionalOptions_2 ON SpatialTable (GeometryColumn)
USING GEOMETRY_GRID
WITH (
    GRIDS = (LEVEL_1 = HIGH, LEVEL_2 = MEDIUM),
    SORT_IN_TEMPDB = ON,
    IGNORE_DUP_KEY = OFF,
    STATISTICS_NORECOMPUTE = OFF
);

CREATE SPATIAL INDEX IX_TraditionalOptions_3 ON SpatialTable (GeometryColumn)
USING GEOMETRY_GRID
WITH (
    CELLS_PER_OBJECT = 16,
    PAD_INDEX = OFF,
    FILLFACTOR = 80,
    SORT_IN_TEMPDB = OFF,
    IGNORE_DUP_KEY = ON,
    STATISTICS_NORECOMPUTE = ON,
    STATISTICS_INCREMENTAL = OFF
);

-- Spatial indexes with file group specification
CREATE SPATIAL INDEX IX_FileGroup_1 ON SpatialTable (GeometryColumn)
USING GEOMETRY_AUTO_GRID
ON [PRIMARY];

CREATE SPATIAL INDEX IX_FileGroup_2 ON SpatialTable (GeometryColumn)
USING GEOMETRY_GRID
WITH (BOUNDING_BOX = (0, 0, 100, 100))
ON [SPATIAL_DATA];

CREATE SPATIAL INDEX IX_FileGroup_3 ON SpatialTable (GeometryColumn)
USING GEOMETRY_GRID
WITH (
    GRIDS = (LEVEL_1 = HIGH, LEVEL_2 = MEDIUM),
    CELLS_PER_OBJECT = 32,
    PAD_INDEX = ON
)
ON [INDEXES];

-- Complex spatial indexes combining all features
CREATE SPATIAL INDEX IX_Complex_1 ON dbo.SpatialTable (GeometryColumn)
USING GEOMETRY_GRID
WITH (
    BOUNDING_BOX = (-180, -90, 180, 90),
    GRIDS = (LEVEL_1 = HIGH, LEVEL_2 = HIGH, LEVEL_3 = MEDIUM, LEVEL_4 = LOW),
    CELLS_PER_OBJECT = 64,
    PAD_INDEX = ON,
    FILLFACTOR = 85,
    SORT_IN_TEMPDB = ON,
    IGNORE_DUP_KEY = OFF,
    STATISTICS_NORECOMPUTE = OFF
)
ON [SPATIAL_INDEXES];

CREATE SPATIAL INDEX IX_Complex_2 ON [GIS].[LocationData] ([Coordinates])
USING GEOGRAPHY_GRID
WITH (
    GRIDS = (LEVEL_1 = HIGH, LEVEL_2 = HIGH, LEVEL_3 = HIGH, LEVEL_4 = MEDIUM),
    CELLS_PER_OBJECT = 128,
    PAD_INDEX = OFF,
    FILLFACTOR = 95,
    SORT_IN_TEMPDB = OFF
)
ON [GEO_DATA];

-- Spatial indexes without semicolon terminator
CREATE SPATIAL INDEX IX_NoSemicolon_1 ON SpatialTable (GeometryColumn) USING GEOMETRY_AUTO_GRID
CREATE SPATIAL INDEX IX_NoSemicolon_2 ON SpatialTable (GeometryColumn) 
USING GEOMETRY_GRID
WITH (BOUNDING_BOX = (0, 0, 100, 100))

-- Spatial indexes with optional semicolon
CREATE SPATIAL INDEX IX_OptionalSemicolon_1 ON SpatialTable (GeometryColumn) USING GEOMETRY_AUTO_GRID;
CREATE SPATIAL INDEX IX_OptionalSemicolon_2 ON SpatialTable (GeometryColumn) 
USING GEOMETRY_GRID
WITH (BOUNDING_BOX = (0, 0, 100, 100));

-- Geography-specific spatial indexes
CREATE SPATIAL INDEX IX_Geography_1 ON GlobalData (LocationColumn) USING GEOGRAPHY_AUTO_GRID;
CREATE SPATIAL INDEX IX_Geography_2 ON GlobalData (LocationColumn) 
USING GEOGRAPHY_GRID
WITH (GRIDS = (LEVEL_1 = HIGH, LEVEL_2 = MEDIUM, LEVEL_3 = LOW));

CREATE SPATIAL INDEX IX_Geography_3 ON GlobalData (LocationColumn) 
USING GEOGRAPHY_GRID
WITH (
    GRIDS = (LEVEL_1 = HIGH, LEVEL_2 = HIGH, LEVEL_3 = MEDIUM, LEVEL_4 = LOW),
    CELLS_PER_OBJECT = 256
);

-- Edge cases and variations
CREATE SPATIAL INDEX IX_EdgeCase_1 ON T (C) USING GEOMETRY_AUTO_GRID;
CREATE SPATIAL INDEX [IX_Brackets] ON [Table] ([Column]) USING GEOMETRY_GRID;
CREATE SPATIAL INDEX IX_MinCells ON SpatialTable (GeometryColumn) 
USING GEOMETRY_GRID 
WITH (CELLS_PER_OBJECT = 1);

CREATE SPATIAL INDEX IX_MaxCells ON SpatialTable (GeometryColumn) 
USING GEOMETRY_GRID 
WITH (CELLS_PER_OBJECT = 8192);

-- Spatial indexes with simple column references only
-- Note: Computed columns/expressions are not supported in spatial index column specifications

-- Multiple spatial indexes on same table
CREATE SPATIAL INDEX IX_Multi_1 ON SpatialTable (GeometryColumn) USING GEOMETRY_AUTO_GRID;
CREATE SPATIAL INDEX IX_Multi_2 ON SpatialTable (GeographyColumn) USING GEOGRAPHY_AUTO_GRID;
CREATE SPATIAL INDEX IX_Multi_3 ON SpatialTable (AnotherGeometryColumn) 
USING GEOMETRY_GRID 
WITH (BOUNDING_BOX = (0, 0, 1000, 1000));