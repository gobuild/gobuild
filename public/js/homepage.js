/* angular */

angular.module('ngApp', ['ngNotify'])
	.filter('fromNow', function(){
		return function(input){
			return moment(input).fromNow();
		}
	})
	.controller('MainCtrl', function($scope, $http, ngNotify){
		$scope.name = 'world'
		$scope.repos = [];

		$http({
			method: "GET",
			url: "/api/recent/repos"
		}).then(function(res){
			$scope.repos = res.data || [];
		}, function(err){
			ngNotify.set(err.data, "error");
		})
	})