/* angular */

angular.module('ngApp', ['ngNotify'])
	.filter('fromNow', function(){
		return function(input){
			return moment(input).fromNow();
		}
	})
	.controller('MainCtrl', function($scope, $http, ngNotify){
		$scope.repos = [];
		$scope.fullname = ''

		var loadRecent = function(){
			return $http({
				method: "GET",
				url: "/api/recent/repos"
			}).then(function(res){
				$scope.repos = res.data || [];
			}, function(err){
				ngNotify.set(err.data, "error");
			})
		}

		loadRecent()

		$scope.addRepository = function(fullname){
			var parts = fullname.split('/')
			if (parts.length != 2){
				ngNotify.set('Format shoule be <owner>/<reop>', 'error')
				return
			}
			ngNotify.set("Checking if repository is valid")

			$http({
				method: "POST",
				url: "/api/repos",
				data: $.param({owner: parts[0], repo: parts[1]}),
				headers: {'Content-Type': 'application/x-www-form-urlencoded'}
			}).then(function(res){
				ngNotify.set(res.data.message);
				loadRecent();
				$scope.fullname = '';
			}, function(err){
				ngNotify.set(err.data, 'error');
			})
		}
	})