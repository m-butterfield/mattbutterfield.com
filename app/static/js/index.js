var TRANSITION_TIME = 200;

var Post = Backbone.Model.extend({
    urlRoot: "/api/post"
});

var PageButtonsView = Backbone.View.extend({
    el: '#page-buttons',

    events: {
        "click .btn": "page"
    },

    template: _.template($('#page-buttons-template').html()),

    initialize: function(options) {
        this.router = options.router;
        this.listenTo(this.model, 'sync', this.render);
    },

    page: function(event) {
        event.preventDefault();
        this.router.navigate(event.target.getAttribute('href'), {trigger: true});
    },

    updatePost: function(postId) {
        this.model.set('id', postId);
        this.model.fetch({reset: true});
    },

    render: function() {
        this.$el.html(this.template(this.model.attributes));
    }
});

var PostContentView = Backbone.View.extend({
    el: '#post-content',

    template: _.template($('#post-content-template').html()),

    initialize: function(options) {
        this.listenTo(this.model, 'sync', this.render);
    },

    render: function() {
        if (this.rendered) {
            this.transitionRender();
        } else {
            this.initialRender();
        }
    },

    transitionRender: function() {
        var that = this;
        this.$el.fadeOut({
            duration: TRANSITION_TIME,
            complete: function() {
                that.$el.html(that.template(that.model.attributes));
                that.$el.fadeIn({duration: TRANSITION_TIME});
            }
        });
    },

    initialRender: function() {
        this.$el.html(this.template(this.model.attributes));
        this.$el.fadeIn({duration: TRANSITION_TIME});
        this.rendered = true;
    }
});

var Router = Backbone.Router.extend({
    routes: {
        "post/:postId": 'loadPost',
        "": "renderViews"
    },

    initialize: function(options) {
        var post = new Post(options.postData);
        this.pageButtonsView = new PageButtonsView({
            model: post,
            router: this
        });
        this.postContentView = new PostContentView({
            model: post
        });
    },

    loadPost: function(postId) {
        if (postId !== this.pageButtonsView.model.id) {
            this.pageButtonsView.updatePost(postId);
        } else {
            this.renderViews();
        }
    },

    renderViews: function() {
        this.pageButtonsView.render();
        this.postContentView.render();
    }
});
