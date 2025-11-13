-- Regression test: Complex multi-table JOINs with ORDER BY
-- Tests parser handling of multiple LEFT JOINs, aliases, and ordering

SELECT
  se.id,
  se.name,
  se.description,
  st.type AS step_type,
  st.index AS step_index,
  sr.conditional,
  sr.index AS sr_index,
  script.script_content
FROM test se
LEFT JOIN test_step st ON se.id = st.sequence_id
LEFT JOIN test_rel sr ON sr.step_id = sr.script_id
LEFT JOIN msg_script script ON script.id = sr.script_id
WHERE se.creator_id = 1
ORDER BY se.id DESC, sr_index ASC;
