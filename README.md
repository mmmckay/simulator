# Simulator

1. Create the layout for the simulation environment
    - [x] Establish coordinate system
    - [x] Grid layout hexagon
    - [ ] Ability to store arbitrary data on grid tiles (and on boundaries?)
    - [ ] Methods for 
        - [ ] Iterating the grid tiles
        - [ ] Selecting a tile/tiles based on coordinates or search criteria
        - [ ] Selecting adjacent/neighboring tiles
    - [x] Way to display grid (simple text is a good start)

2. [ ] Generate terrain and/or environment for the grid
    - [ ] Randomized in some way
    - [ ] Basic implementation will have land/water
    - [ ] Add additional components around biome/land type/elevation
    - Inspiration here https://www.redblobgames.com/maps/terrain-from-noise http://www-cs-students.stanford.edu/~amitp/game-programming/polygon-map-generation/

3. [ ] Flora/Fauna placement
    - [ ] Basic implementation will have plants, prey, and predator entities on the grid
    - [ ] Can be more complex down the line
    - [ ] Define interaction between entities based on criteria

4. [ ] Time System
    - [x] Methods to implement time passage on the grid
    - [ ] Execution order for grid/entities is defined and can be processed

5. [x] Display system from part 1 can handle terrain/entities/time
