import click
import pygame
import random
from pygame.locals import *

from OpenGL.GL import *
from OpenGL.GLU import *

vertices = (
    ( 1, -1, -1), ( 1,  1, -1), (-1,  1, -1), (-1, -1, -1),
    ( 1, -1,  1), ( 1,  1,  1), (-1, -1,  1), (-1,  1,  1)
)
edges = ((0,1),(0,3),(0,4),(2,1),(2,3),(2,7),(6,3),(6,4),(6,7),(5,1),(5,4),(5,7))
surfaces = ((0, 1, 2, 3), (3, 2, 7, 6), (6, 7, 5, 4), (4, 5, 1, 0), (1, 5, 7, 2), (4, 0, 3, 6))
colors = ((1, 0, 0), (0, 1, 0), (1, 0.5, 0), (1, 1, 0), (1, 1, 1), (0, 0, 1))

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

        glBegin(GL_QUADS)
        for i in range(len(surf)):
            glColor3fv(colors[i])
            for j in surf[i]:
                glVertex3fv(vertices[j])
        glEnd()

        glPopMatrix()

class EntireCube():
    def __init__(self, N, scale):
        self.N = N
        cr = range(self.N)
        self.cubes = [Cube((x, y, z), self.N, scale) for x in cr for y in cr for z in cr]

    def mainloop(self, mix):

        rot_cube_map  = { K_UP: (-1, 0), K_DOWN: (1, 0), K_LEFT: (0, -1), K_RIGHT: (0, 1)}
        rot_slice_map = {
            K_l: (0, 0, 1), K_r: (0, 2, 1), K_d: (1, 0, 1),K_u: (1, 2, 1), K_b: (2, 0, 1), K_f: (2, 2, 1),
        }  
        rot_slice_map_prime = {
            K_l: (0, 0, -1), K_r: (0, 2, -1), K_d: (1, 0, -1), K_u: (1, 2, -1), K_b: (2, 0, -1), K_f: (2, 2, -1),
        }
        ang_x, ang_y, rot_cube = 0, 0, (0, 0)
        animate_rot, animate, animate_rot_ang, animate_ang, animate_speed = False, False, 0, 0, 5
        action = (0, 0, 0)

        counter = 0
        arg = ""
        while True:
            
            if not animate and len(mix):

                if not "2" in arg:
                    print("Step %d : %s" % (counter, mix[0]))

                curr = mix[0]
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
                    mix[0] = curr[0] + arg.replace("2", "")
                else:
                    mix.pop(0)
                    counter += 1

            else:
                for event in pygame.event.get():
                    if event.type == pygame.QUIT:
                        pygame.quit()
                        quit()
                    if event.type == KEYDOWN:
                        if not animate_rot and event.key in rot_cube_map:
                            animate_rot, rot_cube = True, rot_cube_map[event.key]
                        if not animate and event.key in rot_slice_map and pygame.key.get_mods() & KMOD_CTRL:
                            animate, action = True, rot_slice_map[event.key]
                        elif not animate and event.key in rot_slice_map_prime:
                            animate, action = True, rot_slice_map_prime[event.key]
                    """if event.type == KEYUP:
                        if event.key in rot_cube_map:
                            rot_cube = (0, 0)"""

            if animate_rot:
                ang_x += rot_cube[0]*animate_speed
                ang_y += rot_cube[1]*animate_speed
                if ang_x % 90 == 0 and ang_y % 90 == 0:
                    animate_rot, animate_rot_ang = False, 0

            glMatrixMode(GL_MODELVIEW)
            glLoadIdentity()
            glTranslatef(0, 0, -40)
            glRotatef(ang_y, 0, 1, 0)
            glRotatef(ang_x, 1, 0, 0)

            glClear(GL_COLOR_BUFFER_BIT|GL_DEPTH_BUFFER_BIT)

            if animate:
                if animate_ang >= 90:
                    for cube in self.cubes:
                        cube.update(*action)
                    animate, animate_ang = False, 0

            for cube in self.cubes:
                cube.draw(colors, surfaces, vertices, animate, animate_ang, *action)
            if animate:
                animate_ang += animate_speed

            pygame.display.flip()
            pygame.time.wait(10)

def parse_mix(mix):
    mix_list = []
    for step in mix.split(" "):
        if len(step) > 0 and len(step) <= 3 and step[0] in "FRUBLD":
            if len(step) >= 2 and len(step) <= 3:
                if len(step) == 2 and step[1] in "'’2":
                    mix_list.append(step)
                elif len(step) == 3 and step[1] in "'’2" and step[2] in "'’2":
                    mix_list.append(step)
                else:
                    print("Error : Invalid step arg")
                    return []
            else:
                mix_list.append(step)
        else:
            print("Error : Invalid step name")
            return []
    return mix_list

@click.command()
@click.argument("mix", default="")
def main(mix):

    if len(mix):
        mix = parse_mix(mix)
    else:
        mix = []

    pygame.init()
    display = (800,600)
    pygame.display.set_mode(display, DOUBLEBUF|OPENGL)
    glEnable(GL_DEPTH_TEST) 

    glMatrixMode(GL_PROJECTION)
    gluPerspective(45, (display[0]/display[1]), 0.1, 50.0)

    glTranslatef(-15, -15, 0)
    glRotatef(25, 1, 0, 0)
    glRotatef(-25, 0, 1, 0)

    NewEntireCube = EntireCube(3, 1.5) 
    NewEntireCube.mainloop(mix)

if __name__ == '__main__':
    main()
    pygame.quit()
    quit()