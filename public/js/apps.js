/* angular */

angular.module('ngApp', [])
	.controller('MainCtrl', function($scope, $http){
		$scope.name = 'world'
		$scope.repos = {}
		$http.get('/api/apps')
			.success(function(res){
				$scope.repos = res;
			})
	})