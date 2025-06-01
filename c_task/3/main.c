#include <gtk/gtk.h>

static void activate(GtkApplication* app, gpointer user_data) {
    GtkWidget *window;
    GtkWidget *vbox;
    GtkWidget *hbox;
    GtkWidget *button1, *button2, *button3;
    
    window = gtk_application_window_new(app);
    gtk_window_set_title(GTK_WINDOW(window), "GTK4 Layout Example");
    gtk_window_set_default_size(GTK_WINDOW(window), 400, 300);
    
    // Create vertical box
    vbox = gtk_box_new(GTK_ORIENTATION_VERTICAL, 5);
    gtk_widget_set_margin_start(vbox, 10);
    gtk_widget_set_margin_end(vbox, 10);
    gtk_widget_set_margin_top(vbox, 10);
    gtk_widget_set_margin_bottom(vbox, 10);
    
    // Create horizontal box
    hbox = gtk_box_new(GTK_ORIENTATION_HORIZONTAL, 5);
    
    // Create buttons
    button1 = gtk_button_new_with_label("Button 1");
    button2 = gtk_button_new_with_label("Button 2");
    button3 = gtk_button_new_with_label("Button 3");
    
    // Add buttons to horizontal box
    gtk_box_append(GTK_BOX(hbox), button1);
    gtk_box_append(GTK_BOX(hbox), button2);
    gtk_box_append(GTK_BOX(hbox), button3);
    
    // Add horizontal box to vertical box
    gtk_box_append(GTK_BOX(vbox), hbox);
    
    gtk_window_set_child(GTK_WINDOW(window), vbox);
    gtk_widget_show(window);
}

int main(int argc, char **argv) {
    GtkApplication *app;
    int status;
    
    app = gtk_application_new("org.example.gtk4", G_APPLICATION_DEFAULT_FLAGS);
    g_signal_connect(app, "activate", G_CALLBACK(activate), NULL);
    status = g_application_run(G_APPLICATION(app), argc, argv);
    g_object_unref(app);
    
    return status;
}
