***

### **Game Design Document: Project Roguelike-Battler (Working Title)**

| **Document Version:** | 1.0 |
| :--- | :--- |
| **Last Updated:** | October 16, 2025 |
| **Authors:** | Walladimus & Vališirdisent |

---

### **Table of Contents**
1.  **High Concept**
2.  **Core Design Pillars**
3.  **Core Gameplay Loop & Inter-phase Mechanics**
4.  **Game Modes**
    4.1. Rogue Mode
    4.2. Story Mode
5.  **Player Progression & Systems**
    5.1. Core Progression (Levels, Gold, Loot)
    5.2. Special In-Game Encounters
    5.3. Item & Upgrade Philosophy
6.  **Social, UI, and Player Experience**
    6.1. Friends System
    6.2. User Interface (UI) & Player Feedback
7.  **Technical Specifications & Monetization**
    7.1. Technical Goals
    7.2. Monetisation
8.  **Development Roadmap & Future Vision**

---

### **1. High Concept**

**Project Rogue-Battle** is a browser-based roguelike action game for solo and multiplayer, blending cooperative PvE progression, competitive PvP battles, and epic Boss encounters into a single, seamless experience. Players will engage in a high-stakes gameplay loop of powering up, proving their skill against others, and teaming up to defeat massive foes. The core player fantasy is the thrill of discovering powerful, synergistic item builds and then using them to outplay opponents in dynamic combat.

---

### **2. Core Design Pillars**

* **Accessible Action:** Browser-based with LAN and online support, designed to be highly performant on a wide range of hardware, including a dedicated "Performance Mode" for low-end systems.
* **Deep Replayability:** A compelling core loop is enhanced by distinct game modes, procedural elements, and a deep item synergy system, ensuring no two runs are the same.
* **Rewarding Challenge:** The difficulty scales intelligently, respecting player skill and strategy. The most demanding challenges yield the greatest rewards and unlock the game's true endings.
* **Meaningful Social Connection:** An integrated friends system encourages players to connect, compete, and cooperate, building a strong community.

---

### **3. Core Gameplay Loop & Inter-phase Mechanics**

The game progresses through a cycle of distinct phases, with performance in one directly influencing the next.

* **The Cycle:** The game alternates primarily between PvE and PvP rounds. After a set number of PvE rounds, a cooperative Boss Round occurs, testing the lobby's collective strength.
* **Inter-phase Bonuses:**
    * **PvE Win -> PvP Advantage:** The top-performing player in a PvE round gains a bonus in the subsequent PvP round: **+5% EXP gain** and **+20% gold from player kills**.
    * **PvE Win -> Boss Advantage:** The top-performing player in a PvE round that precedes a Boss Round starts the boss fight with a **temporary one-hit shield**. This shield is destroyed after the encounter.

---

### **4. Game Modes**

#### **4.1. Rogue Mode**
The core, endlessly replayable experience. The lobby host can define key parameters before the game begins.

* **Competition Mode**
    * **Objective:** To be the last player standing.
    * **Total Health System:** Players begin the run with a fixed pool of Total Health (e.g., 100). Losing in PvP rounds deals damage directly to this pool. A player is eliminated when their Total Health reaches zero.
    * **Match Pacing:** A typical match is designed to last approximately 20 rounds, with a hard cap of 40 rounds.
    * **Total Health Damage Calculation:** Damage taken after a PvP round is weighted:
        * **Placement (50%):** A player's final rank is the most significant factor.
        * **Survival Time (33%):** Players who are eliminated quickly take more damage.
        * **Players Remaining (17%):** Losing early when more players are alive is more punishing.
    * **Comeback Mechanic:** A player who loses Total Health for multiple consecutive rounds or is near elimination early in the game will receive a free compensatory upgrade or ability.

* **Standard Mode**
    * **Objective:** To have the best performance over a pre-set number of rounds.
    * **Gameplay:** A more traditional match structure where players fight in each round, respawning for the next. The overall winner is determined by cumulative performance.
    * **Rewards:** The Top 3 performers in each round receive superior loot.

#### **4.2. Story Mode**
A curated, narrative-driven experience planned for the full release.

* **Objective:** Follow a character's journey through a fixed number of challenging levels, culminating in one of several endings.
* **Features:** Will include unique, story-specific achievements and lore.

---

### **5. Player Progression & Systems**

#### **5.1. Core Progression (Levels, Gold, Loot)**
Players grow stronger by fighting enemies and players, earning EXP to level up, Gold to spend, and Loot to equip. Enemy difficulty and loot quality scale as the run progresses.

#### **5.2. Special In-Game Encounters**
* **The Seller:** An NPC merchant who appears periodically, selling weapons, healing, and upgrades for gold.
* **The Enchanter/Upgrader:** A rare NPC hidden in the environment. Can upgrade a weapon's base stats or enchant it with powerful, unique modifiers. Finding all Enchanters in a single run unlocks a secret ending.

#### **5.3. Item & Upgrade Philosophy**
The item system is the heart of the roguelike experience. Design will prioritise **synergy over linear stat increases**. The goal is to create a rich ecosystem of items and abilities that interact in powerful and unexpected ways, rewarding players who experiment and discover game-breaking combinations.

---

### **6. Social, UI, and Player Experience**

#### **6.1. Friends System**
A robust social hub allowing players to create accounts, add friends by username, view a "recently played with" list, and receive friend suggestions.

#### **6.2. User Interface (UI) & Player Feedback**
* **Main Menu:** A clean, intuitive interface with clear buttons: Play, Achievements, Friends, Settings, Notices & Requests, and a non-intrusive "Buy Me a Coffee" link.
* **In-Game HUD:** Designed for clarity. Temporary bonuses and status effects (like the PvE win bonus) will be displayed as **small icons on the HUD**. Players can **hover over any icon** to get a detailed tooltip explaining the effect and its origin.

---

### **7. Technical Specifications & Monetization**

#### **7.1. Technical Goals**
* **Platform:** Browser-based, playable with no downloads required.
* **Connectivity:** Supports both online multiplayer and local LAN play.
* **Performance:** Highly optimised to run smoothly on low-end hardware. A dedicated graphics setting will toggle between "High Definition" and "Performance (Toaster)" modes.
* **Map Design:** Maps will be large and tactically interesting but built with efficient assets to maintain high performance.

#### **7.2. Monetization**
The game will be free to play. Monetisation will be handled ethically through a **"Buy Me a Coffee"** donation model. Donors will receive unique but purely cosmetic, non-gameplay-altering perks as a thank you.

---

### **8. Development Roadmap & Future Vision**

* **Phase 1: Core Development:** The primary focus is on building and polishing the complete **Rogue Mode** as the foundational pillar of the game. The Story Mode will be developed in parallel.
* **Phase 2: Launch:** The game will launch with **both the Rogue Mode and Story Mode fully implemented.** This ensures a content-rich experience for all players from day one and establishes a stable meta.
* **Post-Launch Vision:** The first major planned expansion is **"Versus Mode,"** a team-based mode designed to support distinct character roles (Tank, Support, etc.) and encourage strategic, cooperative team play.
