-- Spatial Static Method Test Cases
-- Testing GEOGRAPHY::method() and GEOMETRY::method() static method calls

-- Basic geography static methods
SELECT geography::STGeomFromText('POINT(1 2)', 4326);
SELECT geography::STPointFromText('POINT(-122.34900 47.65100)', 4326);
SELECT geography::STLineFromText('LINESTRING(-122.360 47.656, -122.343 47.656)', 4326);
SELECT geography::STPolyFromText('POLYGON((-122.358 47.653, -122.348 47.649, -122.348 47.658, -122.358 47.658, -122.358 47.653))', 4326);
SELECT geography::Point(47.65100, -122.34900, 4326);
SELECT geography::Parse('POINT(-122.34900 47.65100)');

-- Basic geometry static methods
SELECT geometry::STGeomFromText('POINT(1 2)', 0);
SELECT geometry::STGeomFromText('POLYGON((0 0, 1 0, 1 1, 0 1, 0 0))', 0);
SELECT geometry::STPointFromText('POINT(3 4)', 0);
SELECT geometry::STLineFromText('LINESTRING(0 0, 1 1, 2 1, 2 2)', 0);
SELECT geometry::STPolyFromText('POLYGON((0 0, 10 0, 10 10, 0 10, 0 0))', 0);
SELECT geometry::Parse('LINESTRING(0 0, 1 1)');
SELECT geometry::STGeomFromWKB(0x0101000000000000000000F03F0000000000000040, 0);

-- Multi-geometry methods
SELECT geometry::STMPointFromText('MULTIPOINT((1 1), (2 2), (3 3))', 0);
SELECT geometry::STMLineFromText('MULTILINESTRING((0 0, 1 1), (2 2, 3 3))', 0);
SELECT geometry::STMPolyFromText('MULTIPOLYGON(((0 0, 1 0, 1 1, 0 1, 0 0)), ((2 2, 3 2, 3 3, 2 3, 2 2)))', 0);
SELECT geometry::STGeomCollFromText('GEOMETRYCOLLECTION(POINT(1 1), LINESTRING(0 0, 1 1))', 0);

-- Spatial methods in variable assignments
DECLARE @g geography = geography::STGeomFromText('POINT(1 2)', 4326);
DECLARE @point geometry = geometry::Point(3, 4, 0);

-- Spatial methods in INSERT statements
INSERT INTO GeoTable (Location) VALUES (geography::STGeomFromText('POINT(-122.34900 47.65100)', 4326));
INSERT INTO GeoTable (Shape) VALUES (geometry::STGeomFromText('POLYGON((0 0, 10 0, 10 10, 0 10, 0 0))', 0));

-- Spatial methods in UPDATE statements
UPDATE GeoTable SET Location = geography::STGeomFromText('POINT(-122.35 47.65)', 4326) WHERE ID = 1;
UPDATE GeoTable SET Shape = geometry::Parse('POINT(5 5)') WHERE ID = 2;

-- Multiple spatial calls in single statement
SELECT
    geography::Point(47.65, -122.35, 4326) AS Point1,
    geography::Point(47.66, -122.36, 4326) AS Point2,
    geometry::STGeomFromText('POINT(0 0)', 0) AS Origin;

-- Spatial methods in expressions
SELECT 1 WHERE geography::STGeomFromText('POINT(1 2)', 4326) IS NOT NULL;

-- Spatial methods in CASE expressions
SELECT CASE WHEN 1=1 THEN geography::Point(1, 2, 4326) ELSE geography::Point(0, 0, 4326) END;

-- Spatial methods with variables
DECLARE @lat float = 47.65;
DECLARE @lon float = -122.35;
DECLARE @srid int = 4326;
SELECT geography::Point(@lat, @lon, @srid);

-- Spatial methods in subqueries
SELECT * FROM (SELECT geography::Point(1, 2, 4326) AS Location) AS sub;

-- Spatial methods with column references in function arguments
SELECT geography::STGeomFromText(WKTColumn, SRIDColumn) FROM GeoTable;
