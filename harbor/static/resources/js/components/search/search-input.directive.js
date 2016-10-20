/*
    Copyright (c) 2016 VMware, Inc. All Rights Reserved.
    Licensed under the Apache License, Version 2.0 (the "License");
    you may not use this file except in compliance with the License.
    You may obtain a copy of the License at
        
        http://www.apache.org/licenses/LICENSE-2.0
        
    Unless required by applicable law or agreed to in writing, software
    distributed under the License is distributed on an "AS IS" BASIS,
    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
    See the License for the specific language governing permissions and
    limitations under the License.
*/
(function() {
  
  'use strict';
  
  angular
    .module('harbor.search')
    .directive('searchInput', searchInput);
    
  SearchInputController.$inject = ['$scope', '$location', '$window'];
  
  function SearchInputController($scope, $location, $window) {
    var vm = this;

    vm.searchFor = searchFor;
    
    function searchFor(searchContent) {
      $location
        .path('/search')
        .search({'q': searchContent});
      $window.location.href = $location.url();
    }
    
  }
  
  function searchInput() {
    
    var directive = {
      'restrict': 'E',
      'templateUrl': '/static/resources/js/components/search/search-input.directive.html',
      'scope': {
        'searchInput': '=',
      },
      'link': link,
      'controller': SearchInputController,
      'controllerAs': 'vm',
      'bindToController': true
    };
    return directive;
    
    function link(scope, element, attrs, ctrl) {
      element
        .find('input[type="text"]')
        .on('keydown', keydownHandler);
        
      function keydownHandler(e) {
        if($(this).is(':focus') && e.keyCode === 13) {
          ctrl.searchFor($(this).val());
        }
      }
      
    }
  }
  
})();