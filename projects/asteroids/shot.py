import pygame

from circleshape import CircleShape
from constants import SHOT_RADIUS


class Shot(CircleShape):
    def __init__(self, position):
        super().__init__(position.x, position.y, SHOT_RADIUS)

    def draw(self, screen):
        pygame.draw.circle(
            screen, (0, 255, 0), center=self.position, radius=self.radius, width=2
        )

    def update(self, dt):
        self.position += self.velocity * dt
