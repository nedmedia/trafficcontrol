/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

var DeliveryServiceStatsService = function($http, $q, ENV) {

	this.getBPS = function(xmlId, start, end) {
		var request = $q.defer();

		var url = ENV.api['root'] + "deliveryservice_stats",
			params = { deliveryServiceName: xmlId, metricType: 'kbps', serverType: 'edge', startDate: start.seconds(00).format(), endDate: end.seconds(00).format(), interval: '60s' };

		$http.get(url, { params: params })
			.then(
				function(result) {
					request.resolve(result.data.response);
				},
				function(fault) {
					request.reject();
				}
			);

		return request.promise;
	};

	this.getTPS = function(xmlId, start, end) {
		var request = $q.defer();

		var url = ENV.api['root'] + "deliveryservice_stats",
			params = { deliveryServiceName: xmlId, metricType: 'tps_total', serverType: 'edge', startDate: start.seconds(00).format(), endDate: end.seconds(00).format(), interval: '60s' };

		$http.get(url, { params: params })
			.then(
				function(result) {
					request.resolve(result.data.response);
				},
				function(fault) {
					request.reject();
				}
			);

		return request.promise;
	};

};

DeliveryServiceStatsService.$inject = ['$http', '$q', 'ENV'];
module.exports = DeliveryServiceStatsService;
