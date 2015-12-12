/* angular */

angular.module('ngApp', [])
	.controller('MainCtrl', function($scope, $http){
		$scope.repos = {}
		$http({
			method: "GET",
			url: "/api/repos"
		}).then(function(res){
			$scope.repos = res.data;
		}, function(err){
			console.log(err)
		})
	})