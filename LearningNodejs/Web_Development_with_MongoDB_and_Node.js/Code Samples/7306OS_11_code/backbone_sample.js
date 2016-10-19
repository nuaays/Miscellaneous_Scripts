var Person = Backbone.Model.extend();
var PersonView = Backbone.View.extend({
    tag: 'div',
    render: function() {
        var html = [
            this.model.get('name'),
            '<br/>',
            this.model.get('website')
        ].join('');

        this.$el.html(html);

        return this;
    }
});

var person = new Person({
        name: 'Jason Krol',
        website: 'http://kroltech.com'
    }),
    view = new PersonView({ model: person });

$('body').append(view.render().el);
