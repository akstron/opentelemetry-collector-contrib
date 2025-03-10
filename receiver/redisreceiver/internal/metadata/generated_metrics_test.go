// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver/receivertest"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"
)

type testDataSet int

const (
	testDataSetDefault testDataSet = iota
	testDataSetAll
	testDataSetNone
)

func TestMetricsBuilder(t *testing.T) {
	tests := []struct {
		name        string
		metricsSet  testDataSet
		resAttrsSet testDataSet
		expectEmpty bool
	}{
		{
			name: "default",
		},
		{
			name:        "all_set",
			metricsSet:  testDataSetAll,
			resAttrsSet: testDataSetAll,
		},
		{
			name:        "none_set",
			metricsSet:  testDataSetNone,
			resAttrsSet: testDataSetNone,
			expectEmpty: true,
		},
		{
			name:        "filter_set_include",
			resAttrsSet: testDataSetAll,
		},
		{
			name:        "filter_set_exclude",
			resAttrsSet: testDataSetAll,
			expectEmpty: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			start := pcommon.Timestamp(1_000_000_000)
			ts := pcommon.Timestamp(1_000_001_000)
			observedZapCore, observedLogs := observer.New(zap.WarnLevel)
			settings := receivertest.NewNopSettings(receivertest.NopType)
			settings.Logger = zap.New(observedZapCore)
			mb := NewMetricsBuilder(loadMetricsBuilderConfig(t, tt.name), settings, WithStartTime(start))

			expectedWarnings := 0

			assert.Equal(t, expectedWarnings, observedLogs.Len())

			defaultMetricsCount := 0
			allMetricsCount := 0

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordRedisClientsBlockedDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordRedisClientsConnectedDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordRedisClientsMaxInputBufferDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordRedisClientsMaxOutputBufferDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordRedisClusterSlotsAssignedDataPoint(ts, 1)

			allMetricsCount++
			mb.RecordRedisCmdCallsDataPoint(ts, 1, "cmd-val")

			allMetricsCount++
			mb.RecordRedisCmdLatencyDataPoint(ts, 1, "cmd-val", AttributePercentileP50)

			allMetricsCount++
			mb.RecordRedisCmdUsecDataPoint(ts, 1, "cmd-val")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordRedisCommandsDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordRedisCommandsProcessedDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordRedisConnectionsReceivedDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordRedisConnectionsRejectedDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordRedisCPUTimeDataPoint(ts, 1, AttributeStateSys)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordRedisDbAvgTTLDataPoint(ts, 1, "db-val")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordRedisDbExpiresDataPoint(ts, 1, "db-val")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordRedisDbKeysDataPoint(ts, 1, "db-val")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordRedisKeysEvictedDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordRedisKeysExpiredDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordRedisKeyspaceHitsDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordRedisKeyspaceMissesDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordRedisLatestForkDataPoint(ts, 1)

			allMetricsCount++
			mb.RecordRedisMaxmemoryDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordRedisMemoryFragmentationRatioDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordRedisMemoryLuaDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordRedisMemoryPeakDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordRedisMemoryRssDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordRedisMemoryUsedDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordRedisNetInputDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordRedisNetOutputDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordRedisRdbChangesSinceLastSaveDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordRedisReplicationBacklogFirstByteOffsetDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordRedisReplicationOffsetDataPoint(ts, 1)

			allMetricsCount++
			mb.RecordRedisReplicationReplicaOffsetDataPoint(ts, 1)

			allMetricsCount++
			mb.RecordRedisRoleDataPoint(ts, 1, AttributeRoleReplica)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordRedisSlavesConnectedDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordRedisUptimeDataPoint(ts, 1)

			rb := mb.NewResourceBuilder()
			rb.SetRedisVersion("redis.version-val")
			rb.SetServerAddress("server.address-val")
			rb.SetServerPort("server.port-val")
			res := rb.Emit()
			metrics := mb.Emit(WithResource(res))

			if tt.expectEmpty {
				assert.Equal(t, 0, metrics.ResourceMetrics().Len())
				return
			}

			assert.Equal(t, 1, metrics.ResourceMetrics().Len())
			rm := metrics.ResourceMetrics().At(0)
			assert.Equal(t, res, rm.Resource())
			assert.Equal(t, 1, rm.ScopeMetrics().Len())
			ms := rm.ScopeMetrics().At(0).Metrics()
			if tt.metricsSet == testDataSetDefault {
				assert.Equal(t, defaultMetricsCount, ms.Len())
			}
			if tt.metricsSet == testDataSetAll {
				assert.Equal(t, allMetricsCount, ms.Len())
			}
			validatedMetrics := make(map[string]bool)
			for i := 0; i < ms.Len(); i++ {
				switch ms.At(i).Name() {
				case "redis.clients.blocked":
					assert.False(t, validatedMetrics["redis.clients.blocked"], "Found a duplicate in the metrics slice: redis.clients.blocked")
					validatedMetrics["redis.clients.blocked"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Number of clients pending on a blocking call", ms.At(i).Description())
					assert.Equal(t, "{client}", ms.At(i).Unit())
					assert.False(t, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "redis.clients.connected":
					assert.False(t, validatedMetrics["redis.clients.connected"], "Found a duplicate in the metrics slice: redis.clients.connected")
					validatedMetrics["redis.clients.connected"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Number of client connections (excluding connections from replicas)", ms.At(i).Description())
					assert.Equal(t, "{client}", ms.At(i).Unit())
					assert.False(t, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "redis.clients.max_input_buffer":
					assert.False(t, validatedMetrics["redis.clients.max_input_buffer"], "Found a duplicate in the metrics slice: redis.clients.max_input_buffer")
					validatedMetrics["redis.clients.max_input_buffer"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Biggest input buffer among current client connections", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "redis.clients.max_output_buffer":
					assert.False(t, validatedMetrics["redis.clients.max_output_buffer"], "Found a duplicate in the metrics slice: redis.clients.max_output_buffer")
					validatedMetrics["redis.clients.max_output_buffer"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Longest output list among current client connections", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "redis.cluster.slots_assigned":
					assert.False(t, validatedMetrics["redis.cluster.slots_assigned"], "Found a duplicate in the metrics slice: redis.cluster.slots_assigned")
					validatedMetrics["redis.cluster.slots_assigned"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Number of slots which are associated to some node", ms.At(i).Description())
					assert.Equal(t, "{slots}", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "redis.cmd.calls":
					assert.False(t, validatedMetrics["redis.cmd.calls"], "Found a duplicate in the metrics slice: redis.cmd.calls")
					validatedMetrics["redis.cmd.calls"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Total number of calls for a command", ms.At(i).Description())
					assert.Equal(t, "{call}", ms.At(i).Unit())
					assert.True(t, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("cmd")
					assert.True(t, ok)
					assert.EqualValues(t, "cmd-val", attrVal.Str())
				case "redis.cmd.latency":
					assert.False(t, validatedMetrics["redis.cmd.latency"], "Found a duplicate in the metrics slice: redis.cmd.latency")
					validatedMetrics["redis.cmd.latency"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Command execution latency", ms.At(i).Description())
					assert.Equal(t, "s", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.InDelta(t, float64(1), dp.DoubleValue(), 0.01)
					attrVal, ok := dp.Attributes().Get("cmd")
					assert.True(t, ok)
					assert.EqualValues(t, "cmd-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("percentile")
					assert.True(t, ok)
					assert.EqualValues(t, "p50", attrVal.Str())
				case "redis.cmd.usec":
					assert.False(t, validatedMetrics["redis.cmd.usec"], "Found a duplicate in the metrics slice: redis.cmd.usec")
					validatedMetrics["redis.cmd.usec"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Total time for all executions of this command", ms.At(i).Description())
					assert.Equal(t, "us", ms.At(i).Unit())
					assert.True(t, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("cmd")
					assert.True(t, ok)
					assert.EqualValues(t, "cmd-val", attrVal.Str())
				case "redis.commands":
					assert.False(t, validatedMetrics["redis.commands"], "Found a duplicate in the metrics slice: redis.commands")
					validatedMetrics["redis.commands"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Number of commands processed per second", ms.At(i).Description())
					assert.Equal(t, "{ops}/s", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "redis.commands.processed":
					assert.False(t, validatedMetrics["redis.commands.processed"], "Found a duplicate in the metrics slice: redis.commands.processed")
					validatedMetrics["redis.commands.processed"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Total number of commands processed by the server", ms.At(i).Description())
					assert.Equal(t, "{command}", ms.At(i).Unit())
					assert.True(t, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "redis.connections.received":
					assert.False(t, validatedMetrics["redis.connections.received"], "Found a duplicate in the metrics slice: redis.connections.received")
					validatedMetrics["redis.connections.received"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Total number of connections accepted by the server", ms.At(i).Description())
					assert.Equal(t, "{connection}", ms.At(i).Unit())
					assert.True(t, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "redis.connections.rejected":
					assert.False(t, validatedMetrics["redis.connections.rejected"], "Found a duplicate in the metrics slice: redis.connections.rejected")
					validatedMetrics["redis.connections.rejected"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Number of connections rejected because of maxclients limit", ms.At(i).Description())
					assert.Equal(t, "{connection}", ms.At(i).Unit())
					assert.True(t, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "redis.cpu.time":
					assert.False(t, validatedMetrics["redis.cpu.time"], "Found a duplicate in the metrics slice: redis.cpu.time")
					validatedMetrics["redis.cpu.time"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "System CPU consumed by the Redis server in seconds since server start", ms.At(i).Description())
					assert.Equal(t, "s", ms.At(i).Unit())
					assert.True(t, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.InDelta(t, float64(1), dp.DoubleValue(), 0.01)
					attrVal, ok := dp.Attributes().Get("state")
					assert.True(t, ok)
					assert.EqualValues(t, "sys", attrVal.Str())
				case "redis.db.avg_ttl":
					assert.False(t, validatedMetrics["redis.db.avg_ttl"], "Found a duplicate in the metrics slice: redis.db.avg_ttl")
					validatedMetrics["redis.db.avg_ttl"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Average keyspace keys TTL", ms.At(i).Description())
					assert.Equal(t, "ms", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("db")
					assert.True(t, ok)
					assert.EqualValues(t, "db-val", attrVal.Str())
				case "redis.db.expires":
					assert.False(t, validatedMetrics["redis.db.expires"], "Found a duplicate in the metrics slice: redis.db.expires")
					validatedMetrics["redis.db.expires"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Number of keyspace keys with an expiration", ms.At(i).Description())
					assert.Equal(t, "{key}", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("db")
					assert.True(t, ok)
					assert.EqualValues(t, "db-val", attrVal.Str())
				case "redis.db.keys":
					assert.False(t, validatedMetrics["redis.db.keys"], "Found a duplicate in the metrics slice: redis.db.keys")
					validatedMetrics["redis.db.keys"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Number of keyspace keys", ms.At(i).Description())
					assert.Equal(t, "{key}", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("db")
					assert.True(t, ok)
					assert.EqualValues(t, "db-val", attrVal.Str())
				case "redis.keys.evicted":
					assert.False(t, validatedMetrics["redis.keys.evicted"], "Found a duplicate in the metrics slice: redis.keys.evicted")
					validatedMetrics["redis.keys.evicted"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Number of evicted keys due to maxmemory limit", ms.At(i).Description())
					assert.Equal(t, "{key}", ms.At(i).Unit())
					assert.True(t, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "redis.keys.expired":
					assert.False(t, validatedMetrics["redis.keys.expired"], "Found a duplicate in the metrics slice: redis.keys.expired")
					validatedMetrics["redis.keys.expired"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Total number of key expiration events", ms.At(i).Description())
					assert.Equal(t, "{event}", ms.At(i).Unit())
					assert.True(t, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "redis.keyspace.hits":
					assert.False(t, validatedMetrics["redis.keyspace.hits"], "Found a duplicate in the metrics slice: redis.keyspace.hits")
					validatedMetrics["redis.keyspace.hits"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Number of successful lookup of keys in the main dictionary", ms.At(i).Description())
					assert.Equal(t, "{hit}", ms.At(i).Unit())
					assert.True(t, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "redis.keyspace.misses":
					assert.False(t, validatedMetrics["redis.keyspace.misses"], "Found a duplicate in the metrics slice: redis.keyspace.misses")
					validatedMetrics["redis.keyspace.misses"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Number of failed lookup of keys in the main dictionary", ms.At(i).Description())
					assert.Equal(t, "{miss}", ms.At(i).Unit())
					assert.True(t, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "redis.latest_fork":
					assert.False(t, validatedMetrics["redis.latest_fork"], "Found a duplicate in the metrics slice: redis.latest_fork")
					validatedMetrics["redis.latest_fork"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Duration of the latest fork operation in microseconds", ms.At(i).Description())
					assert.Equal(t, "us", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "redis.maxmemory":
					assert.False(t, validatedMetrics["redis.maxmemory"], "Found a duplicate in the metrics slice: redis.maxmemory")
					validatedMetrics["redis.maxmemory"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "The value of the maxmemory configuration directive", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "redis.memory.fragmentation_ratio":
					assert.False(t, validatedMetrics["redis.memory.fragmentation_ratio"], "Found a duplicate in the metrics slice: redis.memory.fragmentation_ratio")
					validatedMetrics["redis.memory.fragmentation_ratio"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Ratio between used_memory_rss and used_memory", ms.At(i).Description())
					assert.Equal(t, "1", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.InDelta(t, float64(1), dp.DoubleValue(), 0.01)
				case "redis.memory.lua":
					assert.False(t, validatedMetrics["redis.memory.lua"], "Found a duplicate in the metrics slice: redis.memory.lua")
					validatedMetrics["redis.memory.lua"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Number of bytes used by the Lua engine", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "redis.memory.peak":
					assert.False(t, validatedMetrics["redis.memory.peak"], "Found a duplicate in the metrics slice: redis.memory.peak")
					validatedMetrics["redis.memory.peak"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Peak memory consumed by Redis (in bytes)", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "redis.memory.rss":
					assert.False(t, validatedMetrics["redis.memory.rss"], "Found a duplicate in the metrics slice: redis.memory.rss")
					validatedMetrics["redis.memory.rss"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Number of bytes that Redis allocated as seen by the operating system", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "redis.memory.used":
					assert.False(t, validatedMetrics["redis.memory.used"], "Found a duplicate in the metrics slice: redis.memory.used")
					validatedMetrics["redis.memory.used"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Total number of bytes allocated by Redis using its allocator", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "redis.net.input":
					assert.False(t, validatedMetrics["redis.net.input"], "Found a duplicate in the metrics slice: redis.net.input")
					validatedMetrics["redis.net.input"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The total number of bytes read from the network", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					assert.True(t, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "redis.net.output":
					assert.False(t, validatedMetrics["redis.net.output"], "Found a duplicate in the metrics slice: redis.net.output")
					validatedMetrics["redis.net.output"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The total number of bytes written to the network", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					assert.True(t, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "redis.rdb.changes_since_last_save":
					assert.False(t, validatedMetrics["redis.rdb.changes_since_last_save"], "Found a duplicate in the metrics slice: redis.rdb.changes_since_last_save")
					validatedMetrics["redis.rdb.changes_since_last_save"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Number of changes since the last dump", ms.At(i).Description())
					assert.Equal(t, "{change}", ms.At(i).Unit())
					assert.False(t, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "redis.replication.backlog_first_byte_offset":
					assert.False(t, validatedMetrics["redis.replication.backlog_first_byte_offset"], "Found a duplicate in the metrics slice: redis.replication.backlog_first_byte_offset")
					validatedMetrics["redis.replication.backlog_first_byte_offset"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "The master offset of the replication backlog buffer", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "redis.replication.offset":
					assert.False(t, validatedMetrics["redis.replication.offset"], "Found a duplicate in the metrics slice: redis.replication.offset")
					validatedMetrics["redis.replication.offset"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "The server's current replication offset", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "redis.replication.replica_offset":
					assert.False(t, validatedMetrics["redis.replication.replica_offset"], "Found a duplicate in the metrics slice: redis.replication.replica_offset")
					validatedMetrics["redis.replication.replica_offset"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Offset for redis replica", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "redis.role":
					assert.False(t, validatedMetrics["redis.role"], "Found a duplicate in the metrics slice: redis.role")
					validatedMetrics["redis.role"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Redis node's role", ms.At(i).Description())
					assert.Equal(t, "{role}", ms.At(i).Unit())
					assert.False(t, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("role")
					assert.True(t, ok)
					assert.EqualValues(t, "replica", attrVal.Str())
				case "redis.slaves.connected":
					assert.False(t, validatedMetrics["redis.slaves.connected"], "Found a duplicate in the metrics slice: redis.slaves.connected")
					validatedMetrics["redis.slaves.connected"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Number of connected replicas", ms.At(i).Description())
					assert.Equal(t, "{replica}", ms.At(i).Unit())
					assert.False(t, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "redis.uptime":
					assert.False(t, validatedMetrics["redis.uptime"], "Found a duplicate in the metrics slice: redis.uptime")
					validatedMetrics["redis.uptime"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Number of seconds since Redis server start", ms.At(i).Description())
					assert.Equal(t, "s", ms.At(i).Unit())
					assert.True(t, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				}
			}
		})
	}
}
