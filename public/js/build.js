/* angular */

angular.module('ngApp', ['ngNotify'])
	.controller('MainCtrl', function($scope, $http, ngNotify){
		$scope.repos = [];
		$scope.message = null;
		$scope.search = '';
		$scope.buildSinceLast = '';
		$scope.isAdmin = false;
		$scope.username = '';
		$scope.email = '';

		var loadRepos = function(){
			$http({
				method: "GET",
				url: "/api/repos"
			}).then(function(res){
				$scope.repos = res.data;
			}, function(err){
				ngNotify.set("Error: " + err.data, 'error')
			})
		}

		loadRepos()
		$scope.refresh = function() {
			$scope.message = "syncing ...";
			$http({
				method: "POST",
				url: "/api/repos"
			}).then(function(res){
				$scope.message = res.data.message;
				loadRepos()
			}, function(err){
				$scope.message = err;
				console.log(err)
			})
		}

		$http({
			method: "GET",
			url: "/api/user"
		}).then(function(res){
			console.log(res.data)
			$scope.username = res.data.name;
			$scope.isAdmin = res.data.admin;
			$scope.email = res.data.email;
			$scope.buildSinceLast = moment(res.data.repo_updated_at).fromNow();
			console.log($scope.buildSinceLast)
		}, function(err){
			console.log(err)
		})

		$scope.build = function(repo){
			console.log(repo.owner, repo.repo)
			$http({
				method: "POST",
				url: "/api/repos/"+repo.id+"/build"
			}).then(function(res){
				ngNotify.set(res.data.message)
			}, function(err){
				ngNotify.set(err.data, 'error')
			})
		}
	})