-- Comprehensive Spatial Index Test Cases
-- This file demonstrates all supported spatial index features and edge cases

-- Table creation with spatial columns for testing
CREATE TABLE [dbo].[SpatialTestTable] (
    [ID] int IDENTITY(1,1) NOT NULL PRIMARY KEY,
    [Name] nvarchar(100),
    [Location] geometry,
    [GlobalPosition] geography,
    [BoundaryShape] geometry,
    [Region] geography
);

-- Basic spatial index variations
CREATE SPATIAL INDEX [IX_Location_AutoGrid] ON [dbo].[SpatialTestTable] ([Location]) USING GEOMETRY_AUTO_GRID;
CREATE SPATIAL INDEX [IX_Location_Grid] ON [dbo].[SpatialTestTable] ([Location]) USING GEOMETRY_GRID;
CREATE SPATIAL INDEX [IX_GlobalPos_AutoGrid] ON [dbo].[SpatialTestTable] ([GlobalPosition]) USING GEOGRAPHY_AUTO_GRID;
CREATE SPATIAL INDEX [IX_GlobalPos_Grid] ON [dbo].[SpatialTestTable] ([GlobalPosition]) USING GEOGRAPHY_GRID;

-- Spatial indexes with bounding box - testing various coordinate ranges
CREATE SPATIAL INDEX [IX_BBox_World] ON [dbo].[SpatialTestTable] ([Location])
USING GEOMETRY_GRID
WITH (BOUNDING_BOX = (-180, -90, 180, 90));

CREATE SPATIAL INDEX [IX_BBox_USA] ON [dbo].[SpatialTestTable] ([BoundaryShape])
USING GEOMETRY_GRID
WITH (BOUNDING_BOX = (-125, 25, -65, 50));

CREATE SPATIAL INDEX [IX_BBox_Custom] ON [dbo].[SpatialTestTable] ([Location])
USING GEOMETRY_GRID
WITH (BOUNDING_BOX = (-10.5, -5.25, 15.75, 20.125));

-- Spatial indexes with grid level configurations
CREATE SPATIAL INDEX [IX_Grid_SingleLevel] ON [dbo].[SpatialTestTable] ([Location])
USING GEOMETRY_GRID
WITH (GRIDS = (LEVEL_1 = HIGH));

CREATE SPATIAL INDEX [IX_Grid_TwoLevels] ON [dbo].[SpatialTestTable] ([Location])
USING GEOMETRY_GRID
WITH (GRIDS = (LEVEL_1 = HIGH, LEVEL_2 = MEDIUM));

CREATE SPATIAL INDEX [IX_Grid_ThreeLevels] ON [dbo].[SpatialTestTable] ([Location])
USING GEOMETRY_GRID
WITH (GRIDS = (LEVEL_1 = HIGH, LEVEL_2 = HIGH, LEVEL_3 = LOW));

CREATE SPATIAL INDEX [IX_Grid_AllLevels] ON [dbo].[SpatialTestTable] ([BoundaryShape])
USING GEOMETRY_GRID
WITH (GRIDS = (LEVEL_1 = HIGH, LEVEL_2 = HIGH, LEVEL_3 = MEDIUM, LEVEL_4 = LOW));

-- Spatial indexes with cells per object variations
CREATE SPATIAL INDEX [IX_Cells_Min] ON [dbo].[SpatialTestTable] ([Location])
USING GEOMETRY_GRID
WITH (CELLS_PER_OBJECT = 1);

CREATE SPATIAL INDEX [IX_Cells_Default] ON [dbo].[SpatialTestTable] ([Location])
USING GEOMETRY_GRID
WITH (CELLS_PER_OBJECT = 16);

CREATE SPATIAL INDEX [IX_Cells_High] ON [dbo].[SpatialTestTable] ([Location])
USING GEOMETRY_GRID
WITH (CELLS_PER_OBJECT = 256);

CREATE SPATIAL INDEX [IX_Cells_Max] ON [dbo].[SpatialTestTable] ([Location])
USING GEOMETRY_GRID
WITH (CELLS_PER_OBJECT = 8192);

-- Complex spatial indexes combining multiple options
CREATE SPATIAL INDEX [IX_Complex_Geometry] ON [dbo].[SpatialTestTable] ([Location])
USING GEOMETRY_GRID
WITH (
    BOUNDING_BOX = (0, 0, 1000, 1000),
    GRIDS = (LEVEL_1 = HIGH, LEVEL_2 = HIGH, LEVEL_3 = MEDIUM, LEVEL_4 = LOW),
    CELLS_PER_OBJECT = 64
);

CREATE SPATIAL INDEX [IX_Complex_Geography] ON [dbo].[SpatialTestTable] ([GlobalPosition])
USING GEOGRAPHY_GRID
WITH (
    GRIDS = (LEVEL_1 = HIGH, LEVEL_2 = HIGH, LEVEL_3 = HIGH, LEVEL_4 = MEDIUM),
    CELLS_PER_OBJECT = 128
);

-- Spatial indexes with traditional index options
CREATE SPATIAL INDEX [IX_IndexOptions_Basic] ON [dbo].[SpatialTestTable] ([Location])
USING GEOMETRY_GRID
WITH (
    BOUNDING_BOX = (-100, -100, 100, 100),
    PAD_INDEX = ON,
    FILLFACTOR = 90
);

CREATE SPATIAL INDEX [IX_IndexOptions_Performance] ON [dbo].[SpatialTestTable] ([BoundaryShape])
USING GEOMETRY_GRID
WITH (
    GRIDS = (LEVEL_1 = HIGH, LEVEL_2 = MEDIUM),
    CELLS_PER_OBJECT = 32,
    PAD_INDEX = OFF,
    FILLFACTOR = 80,
    SORT_IN_TEMPDB = ON,
    IGNORE_DUP_KEY = OFF,
    STATISTICS_NORECOMPUTE = OFF
);

CREATE SPATIAL INDEX [IX_IndexOptions_Advanced] ON [dbo].[SpatialTestTable] ([Location])
USING GEOMETRY_GRID
WITH (
    BOUNDING_BOX = (-50, -50, 50, 50),
    GRIDS = (LEVEL_1 = HIGH, LEVEL_2 = HIGH, LEVEL_3 = MEDIUM, LEVEL_4 = LOW),
    CELLS_PER_OBJECT = 64,
    PAD_INDEX = ON,
    FILLFACTOR = 85,
    SORT_IN_TEMPDB = OFF,
    IGNORE_DUP_KEY = ON,
    STATISTICS_NORECOMPUTE = ON,
    STATISTICS_INCREMENTAL = OFF
);

-- Spatial indexes with file group specifications
CREATE SPATIAL INDEX [IX_FileGroup_Primary] ON [dbo].[SpatialTestTable] ([Location])
USING GEOMETRY_AUTO_GRID
ON [PRIMARY];

CREATE SPATIAL INDEX [IX_FileGroup_Custom] ON [dbo].[SpatialTestTable] ([GlobalPosition])
USING GEOGRAPHY_GRID
WITH (
    GRIDS = (LEVEL_1 = HIGH, LEVEL_2 = MEDIUM),
    CELLS_PER_OBJECT = 32
)
ON [SPATIAL_DATA];

CREATE SPATIAL INDEX [IX_FileGroup_Complex] ON [dbo].[SpatialTestTable] ([BoundaryShape])
USING GEOMETRY_GRID
WITH (
    BOUNDING_BOX = (0, 0, 1000, 1000),
    GRIDS = (LEVEL_1 = HIGH, LEVEL_2 = HIGH, LEVEL_3 = MEDIUM, LEVEL_4 = LOW),
    CELLS_PER_OBJECT = 128,
    PAD_INDEX = ON,
    FILLFACTOR = 90,
    SORT_IN_TEMPDB = ON
)
ON [INDEXES];

-- Schema-qualified table examples
CREATE SPATIAL INDEX [IX_Schema_Simple] ON [GIS].[Locations] ([Coordinates]) USING GEOGRAPHY_AUTO_GRID;

CREATE SPATIAL INDEX [IX_Schema_Complex] ON [Mapping].[SpatialData] ([GeometryField])
USING GEOMETRY_GRID
WITH (
    BOUNDING_BOX = (-1000, -1000, 1000, 1000),
    GRIDS = (LEVEL_1 = HIGH, LEVEL_2 = HIGH, LEVEL_3 = MEDIUM),
    CELLS_PER_OBJECT = 64,
    PAD_INDEX = ON,
    FILLFACTOR = 85
)
ON [GIS_INDEXES];

-- Edge cases and variations
CREATE SPATIAL INDEX [IX_Edge_ShortName] ON [T] ([C]) USING GEOMETRY_AUTO_GRID;
CREATE SPATIAL INDEX [IX_Edge_LongTableName] ON [VeryLongDatabaseName].[VeryLongSchemaName].[VeryLongTableNameForTesting] ([SpatialColumn]) USING GEOGRAPHY_AUTO_GRID;

-- Different grid density combinations
CREATE SPATIAL INDEX [IX_Grid_AllLow] ON [dbo].[SpatialTestTable] ([Location])
USING GEOMETRY_GRID
WITH (GRIDS = (LEVEL_1 = LOW, LEVEL_2 = LOW, LEVEL_3 = LOW, LEVEL_4 = LOW));

CREATE SPATIAL INDEX [IX_Grid_AllMedium] ON [dbo].[SpatialTestTable] ([Location])
USING GEOMETRY_GRID
WITH (GRIDS = (LEVEL_1 = MEDIUM, LEVEL_2 = MEDIUM, LEVEL_3 = MEDIUM, LEVEL_4 = MEDIUM));

CREATE SPATIAL INDEX [IX_Grid_AllHigh] ON [dbo].[SpatialTestTable] ([Location])
USING GEOMETRY_GRID
WITH (GRIDS = (LEVEL_1 = HIGH, LEVEL_2 = HIGH, LEVEL_3 = HIGH, LEVEL_4 = HIGH));

CREATE SPATIAL INDEX [IX_Grid_Mixed] ON [dbo].[SpatialTestTable] ([Location])
USING GEOMETRY_GRID
WITH (GRIDS = (LEVEL_1 = LOW, LEVEL_2 = HIGH, LEVEL_3 = MEDIUM, LEVEL_4 = HIGH));

-- Geography-specific examples with different configurations
CREATE SPATIAL INDEX [IX_Geography_WorldWide] ON [dbo].[SpatialTestTable] ([GlobalPosition])
USING GEOGRAPHY_GRID
WITH (
    GRIDS = (LEVEL_1 = HIGH, LEVEL_2 = HIGH, LEVEL_3 = MEDIUM, LEVEL_4 = LOW),
    CELLS_PER_OBJECT = 512
);

CREATE SPATIAL INDEX [IX_Geography_Regional] ON [dbo].[SpatialTestTable] ([Region])
USING GEOGRAPHY_AUTO_GRID
WITH (CELLS_PER_OBJECT = 64);

-- Test cases without optional semicolons
CREATE SPATIAL INDEX [IX_NoSemicolon1] ON [dbo].[SpatialTestTable] ([Location]) USING GEOMETRY_AUTO_GRID
CREATE SPATIAL INDEX [IX_NoSemicolon2] ON [dbo].[SpatialTestTable] ([GlobalPosition]) 
USING GEOGRAPHY_GRID
WITH (GRIDS = (LEVEL_1 = HIGH, LEVEL_2 = MEDIUM))

-- Test cases with optional semicolons
CREATE SPATIAL INDEX [IX_WithSemicolon1] ON [dbo].[SpatialTestTable] ([Location]) USING GEOMETRY_AUTO_GRID;
CREATE SPATIAL INDEX [IX_WithSemicolon2] ON [dbo].[SpatialTestTable] ([GlobalPosition]) 
USING GEOGRAPHY_GRID
WITH (GRIDS = (LEVEL_1 = HIGH, LEVEL_2 = MEDIUM));

-- Multiple spatial indexes on the same column with different configurations
CREATE SPATIAL INDEX [IX_Multi_Config1] ON [dbo].[SpatialTestTable] ([Location]) USING GEOMETRY_AUTO_GRID;
CREATE SPATIAL INDEX [IX_Multi_Config2] ON [dbo].[SpatialTestTable] ([Location]) 
USING GEOMETRY_GRID 
WITH (BOUNDING_BOX = (0, 0, 100, 100));
CREATE SPATIAL INDEX [IX_Multi_Config3] ON [dbo].[SpatialTestTable] ([Location]) 
USING GEOMETRY_GRID 
WITH (GRIDS = (LEVEL_1 = HIGH, LEVEL_2 = MEDIUM), CELLS_PER_OBJECT = 32);