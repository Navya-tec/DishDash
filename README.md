# Meal Planner Project

A comprehensive Meal Planner application designed for managing meals in shared accommodations. This project helps users organize meal schedules, minimize food wastage, and make meal planning collaborative. It includes features like customizable meal preferences, dish scheduling, a food opt-out option, and a sneak peek of future meals.

## Table of Contents
- [Features](#features)
- [Project Architecture](#project-architecture)
- [Database Schema](#database-schema)
- [Installation](#installation)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
- [Tech Stack](#tech-stack)

## Features

### 1. **Dish Scheduling**
   - Each flat can set up a customized meal plan for breakfast, lunch, and dinner.
   - Meals are chosen based on user preferences and popularity, with new dishes added daily while excluding recently made meals.

### 2. **Meal Preference Management**
   - Users can select preferred dish groups (e.g., Parathas, Pasta) and set priorities within each dish group.
   - The meal plan is adapted to each user’s preferences, with high-priority dishes appearing more frequently in the schedule.

### 3. **Food Opt-Out System**
   - Users can opt out of meals to help reduce food waste.
   - Opt-out data helps determine the correct amount of food needed for each meal, reducing over-ordering and waste.

### 4. **Sneak Peek of Next Day’s Menu**
   - Users can view the menu planned for the next day, helping them decide if they wish to opt out of a meal.

## Project Architecture

The backend of the Meal Planner project is implemented in Go, using the Gin framework. Data is stored in SQLite, and various APIs manage meal planning, dish preferences, opt-outs, and scheduling.

### Workflow
1. **Initialize Dishes**: The `initializeSchedule` function creates an initial 7-day meal plan if fewer than 7 days of meals have been scheduled.
2. **Daily Scheduling**: Each day, the application creates a new meal plan based on user preferences, recent meals, and dish popularity.
3. **User Preferences**: Each user can set dish preferences, and meal plans adjust based on these preferences.
4. **Opt-Out Management**: Users can update their participation in meals to reduce waste.

## Database Schema

### Tables

1. **global_dish**
   - Stores global dishes available for selection in meal plans.
   - Columns: `id`, `name`, `recipe`, `image_url`, `preparation_time`, `dish_cat_list`

2. **dish**
   - Represents the dishes available within a flat, with scores reflecting preferences.
   - Columns: `id`, `global_dish_id`, `flat_id`, `score`, `dish_cat_list`, `meta`

3. **dish_schedule**
   - Tracks the scheduled meals for each flat.
   - Columns: `id`, `flat_id`, `schedule_date`, `breakfast`, `lunch`, `dinner`

4. **dish_group**
   - Stores various dish groups for categorizing dishes.
   - Columns: `id`, `name`

5. **dish_group_global_dish_mapping**
   - Maps global dishes to dish groups.
   - Columns: `id`, `dish_group_id`, `global_dish_id`

## Installation

1. **Clone the Repository**
   ```bash
   git clone https://github.com/Navya-tec/meal-planner.git
   cd meal-planner
2. **Install Dependencies**
   Ensure Go is installed.
   ```bash
   go mod tidy
4. **Run the Application**
   ```bash
   go run main.go
