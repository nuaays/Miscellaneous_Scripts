var App = Ember.Application.create(),
    movies = [{
        title: "Big Trouble in Little China",
        year: "1986"
    }, {
        title: "Aliens",
        year: "1986"
    }];

App.IndexRoute = Ember.Route.extend({
    model: function() {
        return movies;
    }
});


