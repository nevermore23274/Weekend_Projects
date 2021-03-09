import turtle
import random 

# Set variables
m_Height = 1070 # Keep this a bit lower than actual resolution so you can see title
m_Width = 1920

t = turtle.Pen() # Initializing the turtle 

# Settings for window
win = turtle.Screen() # Initialize screen object to variable "win"
win.setup(m_Width, m_Height) # Set the window to whatever the two variables are
win.bgcolor("black") # Windows background color
win.title("Galaxy Generator") # Window title

def myplanet(red, green, blue):
        t.hideturtle() # Hide the turtle
        t.speed(0) # Set turtle speed to almost instant
        t.color(red, green, blue)
        t.begin_fill() # Begin filling shape
        x = random.randint(5,45) # Chooses a random integer for radius
        t.circle(x) # Draw circle of random radius
        t.end_fill() # Finish filling shape
        t.up() # Pickup pen
        y = random.randint(0,360)
        t.seth(y) # Set heading to random direction
        # t.xcor() is turtle's x; t.ycor() is turtle's y
        if t.xcor() < -600 or t.xcor() > 600:
                t.goto(0,0) # Center
        elif t.ycor() < -200 or t.ycor() > 200:
                t.goto(0,0) # Center
        z = random.randint(75,300) # Range for pen to move
        t.forward(z) # Move random amount of pixels
        t.down() # Pen down
        
def esc():
        turtle.bye() # Exits program
        
for i in range(0, 10):
        a = random.randint(0,100)/100.0
        b = random.randint(0,100)/100.0
        c = random.randint(0,100)/100.0
        myplanet(a, b, c) # Random color to function

turtle.listen() # Listening event for keypress

turtle.onkey(esc, "Escape") # Calls esc function if escape key is pressed
turtle.mainloop() # Keeps program running

# red + green = yellow