package schema

/*

# BitSet Index Format

key = _t/<table_id>/<index_id>/<index_key>
value = <bitset>

# Unique Index Format

key = _t/<table_id>/<index_id>/<index_key>/<primary_key>
value = <null>

# Range Index Format

key = _t/<table_id>/<index_id>/<index_key>/<primary_key>
value = <null>

# Primary Key Format

key = _t/<table_id>/0/<primary_key>
value = <data>

*/
