# coding utf-8

import click
import numpy as np
import os
import random
import subprocess
import time

import pygame
from pygame.locals import *

from OpenGL.GL import *
from OpenGL.GLU import *

vertices = (( 1, -1, -1), ( 1,  1, -1), (-1,  1, -1), (-1, -1, -1), ( 1, -1,  1), ( 1,  1,  1), (-1, -1,  1), (-1,  1,  1))
edges = ((0,1),(0,3),(0,4),(2,1),(2,3),(2,7),(6,3),(6,4),(6,7),(5,1),(5,4),(5,7))
surfaces = ((0, 1, 2, 3), (3, 2, 7, 6), (6, 7, 5, 4), (4, 5, 1, 0), (1, 5, 7, 2), (4, 0, 3, 6))
colors = ((0.8, 0, 0), (0, 0, 0.8), (1, 0.5, 0), (0, 0.8, 0), (1, 1, 1), (1, 1, 0))
moves = ("F", "R", "U", "B", "L", "D", "F'", "R'", "U'", "B'", "L'", "D'","F2", "R2", "U2", "B2", "L2", "D2", "F'2", "R'2", "U'2", "B'2", "L'2", "D'2")

class Cube():
    def __init__(self, id, N, scale):
        self.N = N
        self.scale = scale
        self.init_i = [*id]
        self.current_i = [*id]
        self.rot = [[1 if i==j else 0 for i in range(3)] for j in range(3)]

    def isAffected(self, axis, slice, dir):
        return self.current_i[axis] == slice

    def update(self, axis, slice, dir):
        if not self.isAffected(axis, slice, dir):
            return
        i, j = (axis+1) % 3, (axis+2) % 3
        for k in range(3):
            self.rot[k][i], self.rot[k][j] = -self.rot[k][j]*dir, self.rot[k][i]*dir
        self.current_i[i], self.current_i[j] = (
            self.current_i[j] if dir < 0 else self.N - 1 - self.current_i[j],
            self.current_i[i] if dir > 0 else self.N - 1 - self.current_i[i] )

    def transformMat(self):
        scaleA = [[s*self.scale for s in a] for a in self.rot]  
        scaleT = [(p-(self.N-1)/2)*2.1*self.scale for p in self.current_i] 
        return [*scaleA[0], 0, *scaleA[1], 0, *scaleA[2], 0, *scaleT, 1]

    def draw(self, col, surf, vert, animate, angle, axis, slice, dir):
        glPushMatrix()
        if animate and self.isAffected(axis, slice, dir):
            glRotatef( angle*dir, *[1 if i==axis else 0 for i in range(3)] )
        glMultMatrixf( self.transformMat() )
        
        global face_vao, edge_vao

        # draw faces
        glBindVertexArray( face_vao )
        glDrawArrays( GL_QUADS, 0, 6*4 )
        glBindVertexArray( 0 )

        #draw edges
        glColor3f( 0, 0, 0 )
        glBindVertexArray( edge_vao )
        glDrawElements( GL_LINES, 2*12, GL_UNSIGNED_INT, None )
        glBindVertexArray( 0 )

        glPopMatrix()

class EntireCube():
    def __init__(self, N, scale, steps):
        self.N = N
        cr = range(self.N)
        self.cubes = [Cube((x, y, z), self.N, scale) for x in cr for y in cr for z in cr]
        self.steps = steps
        self.hist = ""
        self.reset = False
        self.solving = False
        if len(self.steps):
            print("\nSuffling...")

    def shuffle(self):
        print("\nSuffling...")
        # build steps
        steps_idx = [random.randint(0, len(moves)-1) for _ in range(20)]
        self.steps = [moves[idx] for idx in steps_idx]
        if len(self.hist) == 0:
            self.reset = True
        self.solving = False

    def solve(self):
        print("\nSolving...")
        args = ("./Rubik.exe", self.hist) 
        popen = subprocess.Popen(args, stdout=subprocess.PIPE)
        popen.wait()
        output = popen.stdout.read().decode()
        if len(output):
            self.steps = parse_steps(str(output).replace("\n", ""))
            if len(self.steps):
                self.hist = ""
                self.reset = True
                self.solving = True
                print("Done\n")
            else:
                time.sleep(0.1)
        else:
            print("Error : No Solution")
            time.sleep(0.1)
        

    def mainloop(self):
        rot_cube_map  = {K_UP: (-1, 0), K_DOWN: (1, 0), K_LEFT: (0, -1), K_RIGHT: (0, 1)}
        rot_slice_map = {K_l: (0, 0, 1), K_r: (0, 2, 1), K_d: (1, 0, 1),K_u: (1, 2, 1), K_b: (2, 0, 1), K_f: (2, 2, 1)}  
        rot_slice_map_prime = {K_l: (0, 0, -1), K_r: (0, 2, -1), K_d: (1, 0, -1), K_u: (1, 2, -1), K_b: (2, 0, -1), K_f: (2, 2, -1)}
        ang_x, ang_y, rot_cube = 0, 0, (0, 0)
        animate_rot, animate, animate_ang, animate_speed = False, False, 0, 5
        action = (0, 0, 0)
        steps_counter = 1
        arg = ""
        curr = ""
        last = ""
        
        while True:
            # Clean screen
            glClear(GL_COLOR_BUFFER_BIT|GL_DEPTH_BUFFER_BIT)

            # reset counter after solved
            if self.reset:
                steps_counter = 1
                self.reset = False

            # handle steps
            if not animate and len(self.steps):
                if not "2" in arg:
                    print("Step %d : %s" % (steps_counter, self.steps[0]))
                curr = self.steps[0]
                arg = ""
                if curr[0] == "F":
                    key = K_f
                elif curr[0] == "R":
                    key = K_r
                elif curr[0] == "U":
                    key = K_u
                elif curr[0] == "B":
                    key = K_b
                elif curr[0] == "L":
                    key = K_l
                elif curr[0] == "D":
                    key = K_d
                if len(curr) >= 2:
                    arg = curr[1:]
                if key in rot_slice_map and "'" in arg:
                    animate, action = True, rot_slice_map[key]
                elif key in rot_slice_map_prime:
                    animate, action = True, rot_slice_map_prime[key]
                if "2" in arg:
                    last = curr
                    self.steps[0] = curr[0] + arg.replace("2", "")
                else:
                    self.steps.pop(0)
                    steps_counter += 1
                    if len(last):
                        curr = last
                    last = ""
                    if self.solving == False:
                        self.hist += curr + " "
            # handle events
            else:
                for event in pygame.event.get():
                    if event.type == pygame.QUIT or (event.type == KEYDOWN and event.key == K_ESCAPE):
                        pygame.quit()
                        print("\nGoodBye !")
                        quit()
                    if event.type == KEYDOWN:
                        curr_tab = [" ", " "]
                        edited = False
                        if not animate_rot and event.key in rot_cube_map:
                            animate_rot, rot_cube = True, rot_cube_map[event.key]
                        if not animate and event.key in rot_slice_map and pygame.key.get_mods() & KMOD_CTRL:
                            animate, action = True, rot_slice_map[event.key]
                            curr_tab[1] = "'"
                            edited = True
                        elif not animate and event.key in rot_slice_map_prime:
                            animate, action = True, rot_slice_map_prime[event.key]
                            edited = True
                        if edited:
                            if event.key == K_f:
                                curr_tab[0] = "F"
                            elif event.key == K_r:
                                curr_tab[0] = "R"
                            elif event.key == K_u:
                                curr_tab[0] = "U"
                            elif event.key == K_b:
                                curr_tab[0] = "B"
                            elif event.key == K_l:
                                curr_tab[0] = "L"
                            elif event.key == K_d:
                                curr_tab[0] = "D"
                            if self.solving == True:
                                steps_counter = 1
                                self.solving = False
                                print("\n")
                            print("Step %d : %s" % (steps_counter, "".join(curr_tab)))
                            steps_counter += 1
                            curr = "".join(curr_tab)
                            self.hist += curr + (" " if "'" in curr_tab[1] else "")
            
            # Show buttons
            if len(self.steps) == 0:
                if len(self.hist):
                    button("Solve", -25.4, 7, action=self.solve)
                button("Shuffle", -24.5, 9.5, action=self.shuffle)

            # animate rotations
            if animate_rot:
                ang_x += rot_cube[0]*animate_speed
                ang_y += rot_cube[1]*animate_speed
                if ang_x % 90 == 0 and ang_y % 90 == 0:
                    animate_rot = False
            glMatrixMode(GL_MODELVIEW)
            glLoadIdentity()
            glTranslatef(0, 0, -40)
            glRotatef(ang_y, 0, 1, 0)
            glRotatef(ang_x, 1, 0, 0)
            
            # Print action on screen
            drawText(-1, 8, curr)

            # step animation
            if animate:
                if animate_ang >= 90:
                    for cube in self.cubes:
                        cube.update(*action)
                    animate, animate_ang = False, 0
            for cube in self.cubes:
                cube.draw(colors, surfaces, vertices, animate, animate_ang, *action)
            if animate:
                animate_ang += animate_speed

            # Draw screen
            pygame.display.flip()
            pygame.time.wait(10)

def parse_steps(steps):
    steps_list = steps.split(" ")
    for step in steps_list:
        if len(step) == 0 or len(step) > 3 or step[0] not in "FRUBLD":
            print("Error : Invalid step name")
            return []
        elif len(step) == 2 and step[1] not in "'’2":
            print("Error : Invalid step arg")
            return []
        elif len(step) == 3 and (step[1] not in "'’2" or step[2] not in "2"):
            print("Error : Invalid step arg")
            return []
    return steps_list

def drawText(x, y, textString, fore=(255,255,255,255), back=(0,0,0,255)):
    global display
    font = pygame.font.Font (None, 64)
    textSurface = font.render(textString, True, fore, back)
    textData = pygame.image.tostring(textSurface, "RGBA", True)
    glRasterPos2d(x, y)
    glDrawPixels(textSurface.get_width(), textSurface.get_height(), GL_RGBA, GL_UNSIGNED_BYTE, textData)
    return glGetFloatv(GL_CURRENT_RASTER_POSITION), textSurface.get_width(), textSurface.get_height()

def button(msg,x,y,action=None):
    pos, w, h = drawText(x, y, msg, back=(50,50,50,255))
    x = pos[0]
    y = display[1] - pos[1] - h
    mouse = pygame.mouse.get_pos()
    clicked = pygame.mouse.get_pressed()
    if x+w > mouse[0] > x and y+h > mouse[1] > y:
        if clicked[0] == 1 and action != None:
            action()

@click.command()
@click.argument("steps", default="")
def main(steps):
    if len(steps):
        steps = parse_steps(steps)
    else:
        steps = []

    # Init
    pygame.init()
    global display
    display = (800,600)
    pygame.display.set_mode(display, DOUBLEBUF|OPENGL)

    # Camera
    glMatrixMode(GL_PROJECTION)
    gluPerspective(45, (display[0]/display[1]), 0.1, 50.0)
    # camera position/rotation
    glTranslatef(-20, -15, 0)
    glRotatef(25, 1, 0, 0)
    glRotatef(-30, 0, 1, 0)

    # Modern OpenGL API :

    # enable depth test (less or equal)
    glEnable( GL_DEPTH_TEST )
    glDepthFunc( GL_LEQUAL )

    # enable back face culling (front faces are drawn clockwise)
    glEnable( GL_CULL_FACE )
    glCullFace( GL_BACK )
    glFrontFace( GL_CW )

    global face_vao, edge_vao
    # define the vertex buffers vor the faces
    attribute_array = []
    for face in range(len(surfaces)):
        for vertex in surfaces[face ]:
            attribute_array.append( vertices[vertex] )
            attribute_array.append( colors[face] )

    face_vbos = glGenBuffers(1)
    glBindBuffer(GL_ARRAY_BUFFER, face_vbos)
    glBufferData( GL_ARRAY_BUFFER, np.array( attribute_array, dtype=np.float32 ), GL_STATIC_DRAW )
    glBindBuffer(GL_ARRAY_BUFFER, 0)

    # define the vertex array object for the faces
    face_vao = glGenVertexArrays( 1 )
    glBindVertexArray( face_vao )

    glBindBuffer(GL_ARRAY_BUFFER, face_vbos)
    glVertexPointer( 3, GL_FLOAT, 6*4, None )
    glEnableClientState( GL_VERTEX_ARRAY )  
    glColorPointer( 3, GL_FLOAT, 6*4, ctypes.cast(3*4, ctypes.c_void_p) )
    glEnableClientState( GL_COLOR_ARRAY ) 
    glBindBuffer(GL_ARRAY_BUFFER, 0) 

    glBindVertexArray( 0 )

    # define the vertex buffer for the edges
    edge_vbo = glGenBuffers(1)
    glBindBuffer(GL_ARRAY_BUFFER, edge_vbo)
    glBufferData( GL_ARRAY_BUFFER, np.array( vertices, dtype=np.float32 ), GL_STATIC_DRAW )
    glBindBuffer(GL_ARRAY_BUFFER, 0)

    # define the vertex array object for the edges
    edge_vao = glGenVertexArrays( 1 )
    glBindVertexArray( edge_vao )

    glBindBuffer(GL_ARRAY_BUFFER, edge_vbo)
    glVertexPointer( 3, GL_FLOAT, 0, None ) 
    glEnableClientState( GL_VERTEX_ARRAY ) 
    glBindBuffer(GL_ARRAY_BUFFER, 0) 

    edge_ibo = glGenBuffers(1)
    glBindBuffer( GL_ELEMENT_ARRAY_BUFFER, edge_ibo )
    glBufferData( GL_ELEMENT_ARRAY_BUFFER, np.array( edges, dtype=np.uint32 ), GL_STATIC_DRAW )

    glBindVertexArray( 0 )
    glBindBuffer( GL_ELEMENT_ARRAY_BUFFER, 0 )

    # Build Rubik's Cube and run loop
    NewEntireCube = EntireCube(3, 1.5, steps)
    NewEntireCube.mainloop()

if __name__ == '__main__':
    print("\n" + "#" * 27 + "\n| Actions : | Options :   |\n" + "#" * 27 + \
        "\n| F : Front | ' : reverse |\n| R : Right | 2 : double  |\n| U : Up    |" + \
        "#" * 14 + "\n| B : Back  |\n| L : Left  |\n| D : Down  |\n" + \
        "#" * 13 + "\n")
    # Compile solver
    if not os.path.exists("Rubik") or not os.path.exists("Rubik.exe"):
        args = ("go", "build")
        popen = subprocess.Popen(args)
    # Run
    main()
    pygame.quit()
    quit()