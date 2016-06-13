Cassandra Protobuf
==================

Cassandra protobuf storage implementation in Go.

Methodology
===========

This code maps protobuf types to Cassandra schemas in the following migration-tolerant and intelligent way:

 - Every row has a field "proto" with a binary blob.
 - Row columns are automatically loaded into the proto after deserialization.
 - This allows some fields to be put into the schema, and others to remain binary only.

Design
======

The design of this library is as follows:

 - Accept a protobuf message and a map (string->interface).
 - Fill the protobuf message with the "proto" field if it exists.
 - Fill the remaining fields with the cassandra columns.

The same goes the opposite direction for serialization.

 - Returns a map (string->interface)
 - Accept a protobuf message and a map (string->interface with zero-value)
   - Map serves as a template for a cql row
   - Map MUST contain "proto" field of type "bytes"
   - Map MUST contain at least one other field, presumably used as ID
 - For each field in the map, check for a corresponding proto field
   - If the names match but types do not, REFUSE to serialize (exit with error)
 - Serialize any corresponding field to cassandra field, and zero it in the protobuf
   - proto3 does not store zeroed fields, this will save space
 - Serialize the protobuf to the proto field.

Usage
=====

First, according to your schema, build a template map, for example:

```go
template := make(map[string]interface{})
template["myStringVal"] = ""
template["myIntVal"] = int32(0)
template["myBoolVal"] = false
template["proto"] = make(0, []byte)
```

The column names should match the names of the fields in your proto file.

Any fields in your proto not specified in your template will be serialized in the "proto" field as a binary blob.

You can serialize a protobuf message to a cassandra map for insertion like so:

```go
myMsg := &MyMessage{MyStringVal: "test"}
row, err := marshal.Marshal(myMsg, template)
```

You can then deserialize that row like so:

```go
myMsg := &MyMessage{}
err := marshal.Unmarshal(myMsg, row)
myMsg.MyStringVal == "test" // true
```
