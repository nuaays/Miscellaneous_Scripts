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
    .module('harbor.services.project')
    .factory('ListProjectService', ListProjectService);
  
  ListProjectService.$inject = ['$http', '$log'];
  
  function ListProjectService($http, $log) {
    
    return ListProject;
    
    function ListProject(projectName, isPublic, page, pageSize) {
      $log.info('list project projectName:' + projectName, ', isPublic:' + isPublic);
      return $http
        .get('/api/projects', {
          'params' : {
            'is_public': isPublic,
            'project_name': projectName,
            'page': page,
            'page_size': pageSize 
          }
        });
      
    }
  }
})();